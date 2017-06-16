package routes

import (
	"net/http"
	config "tryhard-platform/config"

	"github.com/labstack/echo"
)

func Router(e *echo.Echo, cfg config.Config) {
	/* favicon.ico */
	e.GET("/favicon.ico", func(c echo.Context) (err error) {
		// No default icon atm
		return c.String(http.StatusOK, "")
	})
	/* Party Routes */
	//e.Get("/party/gen", getPartyGen())
	/* Apps */
	e.GET("/apps", getApps())
	e.GET("/app/:id", getApp())
}
