package adapters

import (
	admin_factories "clinic-api/src/app/admin/factories"
	admin_handlers "clinic-api/src/app/admin/handlers"
	"clinic-api/src/database"
)

type handlers struct {
	Admin admin_handlers.Handler
}

func Init() handlers {
	conn := new(database.DBConf).InitDB()

	return handlers{
		Admin: admin_factories.Factory(conn.DB),
	}
}
