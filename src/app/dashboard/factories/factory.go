package dashboard_factories

import (
	dashboard_handlers "clinic-api/src/app/dashboard/handlers"
	dashboard_repositories "clinic-api/src/app/dashboard/repositories"
	"clinic-api/src/app/dashboard/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) dashboard_handlers.Handler {
	repo := dashboard_repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *dashboard_handlers.NewHandler(serv)
}
