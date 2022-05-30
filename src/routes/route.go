package routes

import (
	"clinic-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()

	route.GET("/", func(ec echo.Context) error {
		status := http.StatusOK
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), map[string]string{
			"data": "server works!",
		})
	})

	return route
}
