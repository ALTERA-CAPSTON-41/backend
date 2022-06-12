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

	// login
	route.POST("/login", handlers.AttemptLoginUser)

	// patient
	route.POST("/patients", handlers.CreatePatientHandler)
	route.GET("/patients", handlers.GetAllPatientsHandler)
	route.GET("/patients/:id", handlers.GetPatientByIDHandler)
	route.PUT("/patients/:id", handlers.EditPatientByIDHandler)
	route.DELETE("/patients/:id", handlers.DeletePatientByIDHandler)

	// polyclinic
	route.POST("/polyclinics", handlers.CreatePolyclinicHandler)
	route.GET("/polyclinics", handlers.GetAllPolyclinicHandler)
	route.GET("/polyclinics/:id", handlers.GetPolyclinicByIDHandler)
	route.PUT("/polyclinics/:id", handlers.EditPolyclinicByIDHandler)
	route.DELETE("/polyclinics/:id", handlers.DeletePolyclinicByIDHandler)

	// doctor
	route.POST("/doctors", handlers.CreateDoctorHandler)
	route.GET("/doctors", handlers.GetAllDoctorsHandler)
	route.GET("/doctors/:id", handlers.GetDoctorByIDHandler)
	route.PUT("/doctors/:id", handlers.EditDoctorByIDHandler)
	route.DELETE("/doctors/:id", handlers.DeleteDoctorByIDHandler)

	// queue
	route.POST("/queues", handlers.CreateQueueHandler)
	route.GET("/queues", handlers.GetAllQueuesHandler)
	route.GET("/queues/:id", handlers.GetQueueByIDHandler)
	route.PUT("/queues/:id", handlers.EditQueueByIDHandler)
	route.DELETE("/queues/:id", handlers.DeleteQueueByIDHandler)

	return route
}
