package polyclinic_factories

import (
	polyclinic_handlers "clinic-api/src/app/polyclinic/handlers"
	polyclinic_repositories "clinic-api/src/app/polyclinic/repositories"
	"clinic-api/src/app/polyclinic/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) polyclinic_handlers.Handler {
	repo := polyclinic_repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *polyclinic_handlers.NewHandler(serv)
}
