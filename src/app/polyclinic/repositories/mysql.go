package repositories

import (
	"clinic-api/src/app/polyclinic"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func (repo *repository) InsertData(data polyclinic.Domain) (int, error) {
	panic("unimplemented")
}

func (repo *repository) SelectAllData() ([]polyclinic.Domain, error) {
	panic("unimplemented")
}

func (repo *repository) SelectDataByID(id int) (*polyclinic.Domain, error) {
	panic("unimplemented")
}

func (repo *repository) UpdateByID(id int, data polyclinic.Domain) error {
	panic("unimplemented")
}

func (repo *repository) DeleteByID(id int) error {
	panic("unimplemented")
}

func NewMySQLRepository(DB *gorm.DB) polyclinic.Repositories {
	return &repository{DB}
}
