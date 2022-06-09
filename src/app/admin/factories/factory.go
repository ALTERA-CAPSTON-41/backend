package admin_factories

import (
	admin_handlers "clinic-api/src/app/admin/handlers"
	admin_repositories "clinic-api/src/app/admin/repositories"
	"clinic-api/src/app/admin/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) admin_handlers.Handler {
	repo := admin_repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *admin_handlers.NewHandler(serv)
}
