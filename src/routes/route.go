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

	// icd10
	route.GET("icd10/:code", caHandler.ICD10.FindICD10ByCodeHandler)

	// defines authenticated routes
	authenticatedRoute := route.Group("", middlewares.VerifyAuthentication())

	patient := authenticatedRoute.Group("/patients")
	patient.POST("", caHandler.Patient.CreatePatientHandler, middlewares.GrantAdmin)
	patient.GET("", caHandler.Patient.HuntPatientByNameOrNIKOrAllHandler)
	patient.GET("/:id", caHandler.Patient.ShowPatientByIDHandler)
	patient.PUT("/:id", caHandler.Patient.AmendPatientByIDHandler, middlewares.GrantAdmin)
	patient.DELETE("/:id", caHandler.Patient.RemovePatientByIDHandler, middlewares.GrantAdmin)

	polyclinic := authenticatedRoute.Group("/polyclinics")
	polyclinic.POST("", caHandler.Polyclinic.CreatePolyclinicHandler, middlewares.GrantAdmin)
	polyclinic.GET("", caHandler.Polyclinic.ShowAllPolyclinicsHandler)
	polyclinic.GET("/:id", caHandler.Polyclinic.ShowPolyclinicByIDHandler)
	polyclinic.PUT("/:id", caHandler.Polyclinic.AmendPolyclinicByIDHandler, middlewares.GrantAdmin)
	polyclinic.DELETE("/:id", caHandler.Polyclinic.RemovePolyclinicByIDHandler, middlewares.GrantAdmin)

	doctor := authenticatedRoute.Group("/doctors")
	doctor.POST("", caHandler.Doctor.CreateDoctorHandler, middlewares.GrantAdmin)
	doctor.GET("", caHandler.Doctor.ShowAllDoctorsHandler)
	doctor.GET("/:id", caHandler.Doctor.ShowDoctorByIDHandler)
	doctor.PUT("/:id", caHandler.Doctor.AmendDoctorByIDHandler, middlewares.GrantAdmin)
	doctor.DELETE("/:id", caHandler.Doctor.RemoveDoctorByIDHandler, middlewares.GrantAdmin)

	nurse := authenticatedRoute.Group("/nurses")
	nurse.POST("", caHandler.Nurse.CreateNurseHandler, middlewares.GrantAdmin)
	nurse.GET("", caHandler.Nurse.ShowAllNursesHandler)
	nurse.GET("/:id", caHandler.Nurse.ShowNurseByIDHandler)
	nurse.PUT("/:id", caHandler.Nurse.AmendNurseByIDHandler, middlewares.GrantAdmin)
	nurse.DELETE("/:id", caHandler.Nurse.RemoveNurseByIDHandler, middlewares.GrantAdmin)

	admin := authenticatedRoute.Group("/admins", middlewares.GrantAdmin)
	admin.POST("", caHandler.Admin.CreateAdminHandler)
	admin.GET("", caHandler.Admin.ShowAllAdminsHandler)
	admin.GET("/:id", caHandler.Admin.ShowAdminByIDHandler)
	admin.PUT("/:id", caHandler.Admin.AmendAdminByIDHandler)
	admin.DELETE("/:id", caHandler.Admin.RemoveAdminByIDHandler)

	queue := authenticatedRoute.Group("/queues")
	queue.POST("", caHandler.Queue.CreateQueueHandler, middlewares.GrantAdmin)
	queue.GET("", caHandler.Queue.ShowAllQueuesHandler)
	queue.PUT("/:id", caHandler.Queue.AmendQueueByIDHandler)
	queue.DELETE("/:id", caHandler.Queue.RemoveQueueByIDHandler, middlewares.GrantAdmin)

	authenticatedRoute.GET("dashboards/:feature", caHandler.Dashboard.ShowTotalHandler)

	// medical record
	medicalRecord := authenticatedRoute.Group("/medical-records")
	medicalRecord.POST("", caHandler.MedicalRecord.CreateMedicalRecordHandler, middlewares.GrantDoctor)
	medicalRecord.GET("/patient/id/:id", caHandler.MedicalRecord.ShowMedicalRecordByPatientIDHandler)
	medicalRecord.GET("/patient/nik/:nik", caHandler.MedicalRecord.ShowMedicalRecordByPatientNIKHandler)
	medicalRecord.GET("/:id", caHandler.MedicalRecord.ShowMedicalRecordByIDHandler)

	// prescription
	prescription := authenticatedRoute.Group("/prescriptions")
	prescription.POST("", caHandler.Presciption.CreatePrescriptionHandler, middlewares.GrantDoctor)
	prescription.GET("/medical-record/:id", caHandler.Presciption.ShowAllPresciptionsByIDHandler)
	prescription.PUT("/:id", caHandler.Presciption.AmendPrescriptionByIDHandler, middlewares.GrantDoctor)
	prescription.DELETE("/:id", caHandler.Presciption.RemovePrescriptionByIDHandler, middlewares.GrantDoctor)

	return route
}
