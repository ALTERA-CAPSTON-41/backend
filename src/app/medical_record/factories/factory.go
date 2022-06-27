package medicalrecord_factories

import (
	medicalrecord_handlers "clinic-api/src/app/medical_record/handlers"
	medicalrecord_repositories "clinic-api/src/app/medical_record/repositories"
	"clinic-api/src/app/medical_record/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) medicalrecord_handlers.Handler {
	repo := medicalrecord_repositories.NewMySQLRepository(conn)
	serv := services.NewServices(repo)
	return *medicalrecord_handlers.NewHandler(serv)
}
