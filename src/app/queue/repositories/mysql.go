package queue_repositories

import (
	"clinic-api/src/app/queue"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// InsertData implements queue.Repositories
func (repo *repository) InserData(data queue.Domain) (string, error) {
	panic("unimplemented")
}

// SelectAllData implements queue.Repositories
func (repo *repository) SelectAllData(polyclinic, from string) ([]queue.Domain, error) {
	panic("unimplemented")
}

// UpdateByID implements queue.Repositories
func (repo *repository) UpdateByID(id string, data queue.Domain) error {
	panic("unimplemented")
}

// DeleteByID implements queue.Repositories
func (repo *repository) DeleteByID(id string) error {
	panic("unimplemented")
}

func NewMySQLRepository(DB *gorm.DB) queue.Repositories {
	return &repository{DB}
}
