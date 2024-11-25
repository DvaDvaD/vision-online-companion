package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/portierglobal/vision-online-companion/api/gen"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (app *application) serve() error {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Prometheus metrics
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
	prometheus.MustRegister(httpRequestsTotal)

	// Middleware to count requests
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Path() != "/metrics" {
				httpRequestsTotal.WithLabelValues(c.Path()).Inc()
			}
			return next(c)
		}
	})

	// Expose metrics endpoint
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Context for graceful shutdown triggered by shutdown endpoint
	shutdownCtx, shutdownCancel := context.WithCancel(context.Background())
	defer shutdownCancel()
	app.shutdownCancel = shutdownCancel

	gen.RegisterHandlers(e, app)

	go func() {
		fmt.Println(Banner)
		fmt.Printf("Listening on :%d\n", Port)
		if err := e.Start(fmt.Sprintf(":%d", Port)); err != nil && err != http.ErrServerClosed {
			app.logger.Fatal().Err(err).Msg("failed to start server")
		}
	}()

	// Context for graceful shutdown triggered by OS signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Select between OS signal and shutdown request
	select {
	case <-ctx.Done():
		// OS signal triggers shutdown
		app.logger.Info().Msg("shutting down server by OS signal")
	case <-shutdownCtx.Done():
		// Shutdown request triggers shutdown
		app.logger.Info().Msg("shutting down server by shutdown request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		return err
	}

	app.logger.Info().Msg("stopped server")

	return nil
}
