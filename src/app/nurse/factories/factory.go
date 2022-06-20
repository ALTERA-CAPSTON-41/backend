package nurse_factories

import (
	nurse_handlers "clinic-api/src/app/nurse/handlers"
	nurse_repositories "clinic-api/src/app/nurse/repositories"
	"clinic-api/src/app/nurse/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) nurse_handlers.Handler {
	repo := nurse_repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *nurse_handlers.NewHandler(serv)
}
