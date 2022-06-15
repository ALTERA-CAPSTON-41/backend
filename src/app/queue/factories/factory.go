package queue_factories

import (
	queue_handlers "clinic-api/src/app/queue/handlers"
	queue_repositories "clinic-api/src/app/queue/repositories"
	"clinic-api/src/app/queue/services"

	"gorm.io/gorm"
)

func Factory(conn *gorm.DB) queue_handlers.Handler {
	repo := queue_repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *queue_handlers.NewHandler(serv)
}
