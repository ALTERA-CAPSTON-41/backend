package prescription_factories

import (
	prescription_handlers "clinic-api/src/app/prescription/handlers"
	prescription_repositories "clinic-api/src/app/prescription/repositories"
	"clinic-api/src/app/prescription/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) prescription_handlers.Handler {
	repo := prescription_repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *prescription_handlers.NewHandler(serv)
}
