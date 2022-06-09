package repositories

import (
	"clinic-api/src/app/admin"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// DeleteByID implements admin.Repositories
func (repo *repository) DeleteByID(id string) (err error) {
	panic("unimplemented")
}

// InsertData implements admin.Repositories
func (repo *repository) InsertData(data admin.Domain) (id string, err error) {
	panic("unimplemented")
}

// SelectAllData implements admin.Repositories
func (repo *repository) SelectAllData() (data []admin.Domain, err error) {
	panic("unimplemented")
}

// SelectDataByID implements admin.Repositories
func (repo *repository) SelectDataByID(id string) (data admin.Domain, err error) {
	panic("unimplemented")
}

// UpdateByID implements admin.Repositories
func (repo *repository) UpdateByID(id string, data admin.Domain) (err error) {
	panic("unimplemented")
}

func NewMySQLRepository(DB *gorm.DB) admin.Repositories {
	return &repository{DB}
}
