package admin_factories

import (
	"clinic-api/src/app/admin/handlers"
	"clinic-api/src/app/admin/repositories"
	"clinic-api/src/app/admin/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) handlers.Handler {
	repo := repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *handlers.NewHandler(serv)
}
