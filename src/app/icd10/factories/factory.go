package icd10_factories

import (
	icd10_handlers "clinic-api/src/app/icd10/handlers"
	icd10_repositories "clinic-api/src/app/icd10/repositories"
	"clinic-api/src/app/icd10/services"
)

func Factory() icd10_handlers.Handler {
	repo := icd10_repositories.NewMySQLRepository()
	serv := services.NewService(repo)
	return *icd10_handlers.NewHandler(serv)
}
