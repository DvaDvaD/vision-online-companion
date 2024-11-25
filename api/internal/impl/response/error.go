package response

import (
	"github.com/labstack/echo/v4"
)

func ErrBadFormat(ctx echo.Context, err error) error {
	return ctx.JSON(
		400,
		map[string]interface{}{
			"status":  "error",
			"message": "Request failed: " + err.Error(),
			"data":    struct{}{},
		},
	)
}

func ErrRouteNotFound(ctx echo.Context) error {
	return ctx.JSON(
		404,
		map[string]interface{}{
			"status":  "error",
			"message": "Route not found",
			"data":    struct{}{},
		},
	)
}
