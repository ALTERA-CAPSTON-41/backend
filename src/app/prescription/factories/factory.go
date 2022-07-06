package prescription_factories

import (
	"clinic-api/src/app/prescription/services"
	prescription_handlers "clinic-api/src/app/prescription/handlers"
	prescription_repositories "clinic-api/src/app/prescription/repositories"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) prescription_handlers.Handler {
	repo := prescription_repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *prescription_handlers.NewHandler(serv)
}