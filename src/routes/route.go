package routes

import (
	"clinic-api/src/adapters"
	"clinic-api/src/handlers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()
	caHandler := adapters.Init()

	route.GET("/", caHandler.APISpec.GetAPISpec)
	route.GET("/attachments/api-spec.yml", caHandler.APISpec.ServeDocsFile)

	// login
	route.POST("/login", caHandler.Account.AttemptLoginHandler)

	// patient
	route.POST("/patients", handlers.CreatePatientHandler)
	route.GET("/patients", handlers.GetAllPatientsHandler)
	route.GET("/patients/:id", handlers.GetPatientByIDHandler)
	route.PUT("/patients/:id", handlers.EditPatientByIDHandler)
	route.DELETE("/patients/:id", handlers.DeletePatientByIDHandler)

	// polyclinic
	route.POST("/polyclinics", caHandler.Polyclinic.CreatePolyclinicHandler)
	route.GET("/polyclinics", caHandler.Polyclinic.ShowAllPolyclinicsHandler)
	route.GET("/polyclinics/:id", caHandler.Polyclinic.ShowPolyclinicByIDHandler)
	route.PUT("/polyclinics/:id", caHandler.Polyclinic.AmendPolyclinicByIDHandler)
	route.DELETE("/polyclinics/:id", caHandler.Polyclinic.RemovePolyclinicByIDHandler)

	// doctor
	route.POST("/doctors", caHandler.Doctor.CreateDoctorHandler)
	route.GET("/doctors", caHandler.Doctor.ShowAllDoctorsHandler)
	route.GET("/doctors/:id", caHandler.Doctor.ShowDoctorByIDHandler)
	route.PUT("/doctors/:id", caHandler.Doctor.AmendDoctorByIDHandler)
	route.DELETE("/doctors/:id", caHandler.Doctor.RemoveDoctorByIDHandler)

	// nurse
	route.POST("/nurses", caHandler.Nurse.CreateNurseHandler)
	route.GET("/nurses", caHandler.Nurse.ShowAllNursesHandler)
	route.GET("/nurses/:id", caHandler.Nurse.ShowNurseByIDHandler)
	route.PUT("/nurses/:id", caHandler.Nurse.AmendNurseByIDHandler)
	route.DELETE("/nurses/:id", caHandler.Nurse.RemoveNurseByIDHandler)

	// admin
	route.POST("/admins", caHandler.Admin.CreateAdminHandler)
	route.GET("/admins", caHandler.Admin.ShowAllAdminsHandler)
	route.GET("/admins/:id", caHandler.Admin.ShowAdminByIDHandler)
	route.PUT("/admins/:id", caHandler.Admin.AmendAdminByIDHandler)
	route.DELETE("/admins/:id", caHandler.Admin.RemoveAdminByIDHandler)

	// queue
	route.POST("/queues", caHandler.Queue.CreateQueueHandler)
	route.GET("/queues", caHandler.Queue.ShowAllQueuesHandler)
	route.PUT("/queues/:id", caHandler.Queue.AmendQueueByIDHandler)
	route.DELETE("/queues/:id", caHandler.Queue.RemoveQueueByIDHandler)

	return route
}
