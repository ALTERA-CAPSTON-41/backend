package adapters

import (
	account_factories "clinic-api/src/app/account/factories"
	account_handlers "clinic-api/src/app/account/handlers"
	admin_factories "clinic-api/src/app/admin/factories"
	admin_handlers "clinic-api/src/app/admin/handlers"
	apispec_factories "clinic-api/src/app/apispec/factories"
	apispec_handlers "clinic-api/src/app/apispec/handlers"
	dashboard_factories "clinic-api/src/app/dashboard/factories"
	dashboard_handlers "clinic-api/src/app/dashboard/handlers"
	doctor_factories "clinic-api/src/app/doctor/factories"
	doctor_handlers "clinic-api/src/app/doctor/handlers"
	icd10_factories "clinic-api/src/app/icd10/factories"
	icd10_handlers "clinic-api/src/app/icd10/handlers"
	medicalrecord_factories "clinic-api/src/app/medical_record/factories"
	medicalrecord_handlers "clinic-api/src/app/medical_record/handlers"
	nurse_factories "clinic-api/src/app/nurse/factories"
	nurse_handlers "clinic-api/src/app/nurse/handlers"
	patient_factories "clinic-api/src/app/patient/factories"
	patient_handlers "clinic-api/src/app/patient/handlers"
	polyclinic_factories "clinic-api/src/app/polyclinic/factories"
	polyclinic_handlers "clinic-api/src/app/polyclinic/handlers"
	prescription_factories "clinic-api/src/app/prescription/factories"
	prescription_handlers "clinic-api/src/app/prescription/handlers"
	queue_factories "clinic-api/src/app/queue/factories"
	queue_handlers "clinic-api/src/app/queue/handlers"
	"clinic-api/src/database"
)

type handlers struct {
	Account       account_handlers.Handler
	Admin         admin_handlers.Handler
	APISpec       apispec_handlers.Handler
	Polyclinic    polyclinic_handlers.Handler
	Queue         queue_handlers.Handler
	Doctor        doctor_handlers.Handler
	Nurse         nurse_handlers.Handler
	Patient       patient_handlers.Handler
	Dashboard     dashboard_handlers.Handler
	MedicalRecord medicalrecord_handlers.Handler
	ICD10         icd10_handlers.Handler
	Presciption   prescription_handlers.Handler
}

func Init() handlers {
	conn := new(database.DBConf).InitDB()

	return handlers{
		Account:       account_factories.Factory(conn.DB),
		Admin:         admin_factories.Factory(conn.DB),
		APISpec:       apispec_factories.Factory(),
		Polyclinic:    polyclinic_factories.Factory(conn.DB),
		Queue:         queue_factories.Factory(conn.DB),
		Doctor:        doctor_factories.Factory(conn.DB),
		Nurse:         nurse_factories.Factory(conn.DB),
		Patient:       patient_factories.Factory(conn.DB),
		Dashboard:     dashboard_factories.Factory(conn.DB),
		MedicalRecord: medicalrecord_factories.Factory(conn.DB),
		ICD10:         icd10_factories.Factory(),
		Presciption:   prescription_factories.Factory(conn.DB),
	}
}
