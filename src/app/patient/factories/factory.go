package patient_factories

import (
	patient_handlers "clinic-api/src/app/patient/handlers"
	patient_repositories "clinic-api/src/app/patient/repositories"
	"clinic-api/src/app/patient/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) patient_handlers.Handler {
	repo := patient_repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *patient_handlers.NewHandler(serv)
}
