package routes

import (
	"clinic-api/src/adapters"
	"clinic-api/src/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()
	caHandler := adapters.Init()

	// docs
	route.GET("/docs", caHandler.APISpec.GetAPISpec)
	route.GET("/attachments/api-spec.yml", caHandler.APISpec.ServeDocsFile)

	// login
	route.POST("/login", caHandler.Account.AttemptLoginHandler)

	// patient
	patient := route.Group("/patients")
	patient.POST("", caHandler.Patient.CreatePatientHandler)
	patient.GET("", caHandler.Patient.HuntPatientByNameOrNIKOrAllHandler)
	patient.GET("/:id", caHandler.Patient.ShowPatientByIDHandler)
	patient.PUT("/:id", caHandler.Patient.AmendPatientByIDHandler)
	patient.DELETE("/:id", caHandler.Patient.RemovePatientByIDHandler)

	// polyclinic
	polyclinic := route.Group("/polyclinics")
	polyclinic.POST("", caHandler.Polyclinic.CreatePolyclinicHandler)
	polyclinic.GET("", caHandler.Polyclinic.ShowAllPolyclinicsHandler)
	polyclinic.GET("/:id", caHandler.Polyclinic.ShowPolyclinicByIDHandler)
	polyclinic.PUT("/:id", caHandler.Polyclinic.AmendPolyclinicByIDHandler)
	polyclinic.DELETE("/:id", caHandler.Polyclinic.RemovePolyclinicByIDHandler)

	// doctor
	doctor := route.Group("/doctors")
	doctor.POST("", caHandler.Doctor.CreateDoctorHandler)
	doctor.GET("", caHandler.Doctor.ShowAllDoctorsHandler)
	doctor.GET("/:id", caHandler.Doctor.ShowDoctorByIDHandler)
	doctor.PUT("/:id", caHandler.Doctor.AmendDoctorByIDHandler)
	doctor.DELETE("/:id", caHandler.Doctor.RemoveDoctorByIDHandler)

	// nurse
	nurse := route.Group("/nurses")
	nurse.POST("", caHandler.Nurse.CreateNurseHandler)
	nurse.GET("", caHandler.Nurse.ShowAllNursesHandler)
	nurse.GET("/:id", caHandler.Nurse.ShowNurseByIDHandler)
	nurse.PUT("/:id", caHandler.Nurse.AmendNurseByIDHandler)
	nurse.DELETE("/:id", caHandler.Nurse.RemoveNurseByIDHandler)

	// admin
	admin := route.Group("/admins")
	admin.POST("", caHandler.Admin.CreateAdminHandler)
	admin.GET("", caHandler.Admin.ShowAllAdminsHandler)
	admin.GET("/:id", caHandler.Admin.ShowAdminByIDHandler)
	admin.PUT("/:id", caHandler.Admin.AmendAdminByIDHandler)
	admin.DELETE("/:id", caHandler.Admin.RemoveAdminByIDHandler)

	// queue
	queue := route.Group("/queues")
	queue.POST("", caHandler.Queue.CreateQueueHandler)
	queue.GET("", caHandler.Queue.ShowAllQueuesHandler)
	queue.PUT("/:id", caHandler.Queue.AmendQueueByIDHandler)
	queue.DELETE("/:id", caHandler.Queue.RemoveQueueByIDHandler)

	// dashboard
	route.GET("dashboards/:feature", caHandler.Dashboard.ShowTotalHandler)

	// icd10
	route.GET("icd10/:code", caHandler.ICD10.FindICD10ByCodeHandler)

	// medical record
	medicalRecord := route.Group("/medical-records", middlewares.VerifyAuthentication())
	medicalRecord.POST("", caHandler.MedicalRecord.CreateMedicalRecordHandler, middlewares.GrantDoctor)
	medicalRecord.GET("/patient/:nik", caHandler.MedicalRecord.ShowMedicalRecordByPatientNIKHandler)
	medicalRecord.GET("/:id", caHandler.MedicalRecord.ShowMedicalRecordByIDHandler)

	return route
}
