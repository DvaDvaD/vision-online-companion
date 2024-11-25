package main

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// UserAgentMiddleware checks if the User-Agent contains "Portier Vision"
func UserAgentMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userAgent := c.Request().Header.Get("X-Portier-Agent")

		// Example: portier/Vision (Windows 11; v5.0.1)
		userAgentPattern := `^portier\/\w+ \(\w+ [\d.]+; v[\d.]+\)$`
		matched, _ := regexp.MatchString(userAgentPattern, userAgent)
		log.Info().Msgf("User-Agent: %s", userAgent)

		if !matched {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "API only available from portier Applications",
			})
		}

		// Continue to the next handler
		return next(c)
	}
}
