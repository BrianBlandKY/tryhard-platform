package routes

import (
    "net/http"
	"github.com/labstack/echo"
	config "dimples-api/config"
)

func Router(e *echo.Echo, cfg config.Config) {
    /* favicon.ico */
    e.Get("/favicon.ico", func (c echo.Context)(err error) {
        // No default icon atm
        return c.String(http.StatusOK, "")
    })
	/* Party Routes */
	//e.Get("/party/gen", getPartyGen())
	/* Apps */
	e.Get("/apps", getApps())
	e.Get("/app/:id", getApp())
}