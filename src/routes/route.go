package routes

import (
	"clinic-api/src/handlers"
	"clinic-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()

	route.GET("/", func(ec echo.Context) error {
		status := http.StatusOK
		return utils.CreateEchoResponse(ec, status, map[string]string{
			"data": "server works!",
		})
	})

	// patient
	route.POST("/patients", handlers.CreatePatientHandler)
	route.GET("/patients", handlers.GetAllPatientsHandler)
	route.GET("/patients/:id", handlers.GetPatientByIDHandler)
	route.PUT("/patients/:id", handlers.EditDoctorByIDHandler)
	route.DELETE("/patients/:id", handlers.DeletePatientByIDHandler)

	return route
}
