package repositories

import (
	"clinic-api/src/app/doctor"
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// DeleteByID implements doctor.Repositories
func (repo *repository) DeleteByID(id string) error {
	query := repo.DB.Where("user_id", id).Delete(new(Doctor))

	if query.RowsAffected <= 0 && query.Error == nil {
		return errors.New("record not found")
	}
	return query.Error
}

// InsertData implements doctor.Repositories
func (repo *repository) InsertData(data doctor.Domain) (string, error) {
	newDoctor := MapToNewRecord(data)
	return newDoctor.User.ID.String(), repo.DB.Create(&newDoctor).Error
}

// SelectAllData implements doctor.Repositories
func (repo *repository) SelectAllData() ([]doctor.Domain, error) {
	var records []Doctor
	err := repo.DB.Preload("User").Preload("Polyclinic").Find(&records).Error
	return MapToBatchDomain(records), err
}

// SelectDataByID implements doctor.Repositories
func (repo *repository) SelectDataByID(id string) (*doctor.Domain, error) {
	var record Doctor
	err := repo.DB.Preload("User").Preload("Polyclinic").Where("user_id", id).Find(&record).Error
	result := MapToDomain(record)
	return &result, err
}

// UpdateByID implements doctor.Repositories
func (repo *repository) UpdateByID(id string, domain doctor.Domain) error {
	data := MapToExistingRecord(domain)
	query := repo.DB.Where("user_id", id).Omit("password", "user_id", "id").Updates(&data)

	if query.RowsAffected <= 0 && query.Error == nil {
		return errors.New("record not found")
	}
	return query.Error
}

func NewMySQLRepository(DB *gorm.DB) doctor.Repositories {
	return &repository{DB}
}