package adapters

import (
	admin_factories "clinic-api/src/app/admin/factories"
	admin_handlers "clinic-api/src/app/admin/handlers"
	apispec_factories "clinic-api/src/app/apispec/factories"
	apispec_handlers "clinic-api/src/app/apispec/handlers"
	doctor_factories "clinic-api/src/app/doctor/factories"
	doctor_handlers "clinic-api/src/app/doctor/handlers"
	polyclinic_factories "clinic-api/src/app/polyclinic/factories"
	polyclinic_handlers "clinic-api/src/app/polyclinic/handlers"
	queue_factories "clinic-api/src/app/queue/factories"
	queue_handlers "clinic-api/src/app/queue/handlers"
	"clinic-api/src/database"
)

type handlers struct {
	Admin      admin_handlers.Handler
	APISpec    apispec_handlers.Handler
	Polyclinic polyclinic_handlers.Handler
	Queue      queue_handlers.Handler
	Doctor     doctor_handlers.Handler
}

func Init() handlers {
	conn := new(database.DBConf).InitDB()

	return handlers{
		Admin:      admin_factories.Factory(conn.DB),
		APISpec:    apispec_factories.Factory(),
		Polyclinic: polyclinic_factories.Factory(conn.DB),
		Queue:      queue_factories.Factory(conn.DB),
		Doctor:     doctor_factories.Factory(conn.DB),
	}
}
