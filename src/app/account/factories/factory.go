package account_factories

import (
	account_handlers "clinic-api/src/app/account/handlers"
	account_repositories "clinic-api/src/app/account/repositories"
	"clinic-api/src/app/account/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) account_handlers.Handler {
	repo := account_repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *account_handlers.NewHandler(serv)
}
