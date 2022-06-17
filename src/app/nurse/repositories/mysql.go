package nurse_repositories

import (
	"clinic-api/src/app/nurse"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// InsertData implements nurse.Repositories
func (repo *repository) InsertData(data nurse.Domain) (string, error) {
	panic("unimplemented")
}

// SelectAllData implements nurse.Repositories
func (repo *repository) SelectAllData() ([]nurse.Domain, error) {
	panic("unimplemented")
}

func (repo *repository) SelectDataByID(id string) (*nurse.Domain, error) {
	panic("unimplemented")
}

func (repo *repository) UpdateByID(id string, domain nurse.Domain) error {
	panic("unimplemented")
}

func (repo *repository) DeleteByID(id string) error {
	panic("unimplemented")
}

func NewMySQLRepository(DB *gorm.DB) nurse.Repositories {
	return &repository{DB}
}
