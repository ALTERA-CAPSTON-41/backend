package doctor_factories

import (
	doctor_handlers "clinic-api/src/app/doctor/handlers"
	doctor_repositories "clinic-api/src/app/doctor/repositories"
	"clinic-api/src/app/doctor/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) doctor_handlers.Handler {
	repo := doctor_repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *doctor_handlers.NewHandler(serv)
}
