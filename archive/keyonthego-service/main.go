package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/portierglobal/keyonthego-service/src/config"
	"github.com/portierglobal/keyonthego-service/src/utils"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

//go:embed docs/*
var embeddedDocs embed.FS

// CustomValidator holds the validator instance
type CustomValidator struct {
	validator *validator.Validate
}

// Validate implements the echo.Validator interface
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	config.InitConfig()
	utils.GetTMPPath()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Register the validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// TODO only for local development
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Create a filesystem from the embedded docs
	docFS, err := fs.Sub(embeddedDocs, "docs")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Shutdown endpoint context
	shutdownCtx, shutdownCancel := context.WithCancel(context.Background())

	// Shutdown route
	e.POST("/shutdown", func(c echo.Context) error {
		go func() {
			// Delay a bit before shutting down so the response can be sent
			time.Sleep(1 * time.Second)
			shutdownCancel() // Trigger server shutdown
		}()

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Server shutting down...",
		})
	})

	signRoute := e.Group("/key-otg/sign")
	signRoute.Use(UserAgentMiddleware)
	signRoute.POST("", CreateSign)
	signRoute.GET("/:requestID", GetSign)
	signRoute.POST("/:requestID", SubmitSign)
	signRoute.GET("/:requestID/qr", GetURLasQR)

	// Serve the embedded docs at /docs
	e.GET("/docs/*", echo.WrapHandler(http.StripPrefix("/docs/", http.FileServer(http.FS(docFS)))))
	e.GET("/docs", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})

	// Context for graceful shutdown triggered by OS signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start server
	go func() {
		fmt.Println(Banner)
		NetworkInterfaces, err := GetIPAddress()
		if err != nil {
			e.Logger.Fatal("failed to get network interfaces:", err)
		}
		for _, network := range NetworkInterfaces {
			mode := viper.GetString("MODE")
			fmt.Printf("Running in %s mode. ", mode)
			if mode == "cloud" {
				fmt.Printf("Listening on cloud https://service.portierglobal.com\n")
			} else {
				fmt.Printf("Listening on %s http://%s:%d\n", network.Name, network.IP, Port)
			}
		}
		if err := e.Start(fmt.Sprintf(":%d", Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Select between OS signal and shutdown request
	select {
	case <-ctx.Done():
		// OS signal triggers shutdown
	case <-shutdownCtx.Done():
		// Shutdown request triggers shutdown
	}

	// Cleanup tmp/requests folder
	if err := os.RemoveAll(viper.GetString("TMP_FOLDER")); err != nil {
		e.Logger.Error("failed to cleanup tmp/requests folder:", err)
	}

	// Gracefully shutdown the server with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
