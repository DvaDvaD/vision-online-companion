package response

import (
	"github.com/labstack/echo/v4"
)

func SuccessOk(ctx echo.Context, data interface{}) error {
	return ctx.JSON(
		200,
		map[string]interface{}{
			"status":  "success",
			"message": "Request fulfilled successfully",
			"data":    data,
		},
	)
}
func SuccessCreated(ctx echo.Context, data interface{}) error {
	return ctx.JSON(
		201,
		map[string]interface{}{
			"status":  "success",
			"message": "Request fulfilled successfully",
			"data":    data,
		},
	)
}
