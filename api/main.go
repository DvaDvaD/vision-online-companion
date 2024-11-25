package main

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"

	kotg_api "github.com/portierglobal/vision-online-companion/api/internal/gen/kotg"
	kotg_impl "github.com/portierglobal/vision-online-companion/api/internal/impl/kotg"
	"github.com/portierglobal/vision-online-companion/api/internal/impl/response"
)

func main() {
	e := echo.New()
	e.RouteNotFound("/*", response.ErrRouteNotFound)

	// kotg routes
	kotg_server := kotg_impl.NewServer()
	kotg_api.RegisterHandlersWithBaseURL(e, kotg_server, "/api/v1")

	// List all
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		e.Logger.Fatal(err)
	}
	os.WriteFile("routes.json", data, 0644)
	e.Logger.Fatal(e.Start(":1323"))
}
