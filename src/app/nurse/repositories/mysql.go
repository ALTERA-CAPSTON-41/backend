package nurse_repositories

import (
	"clinic-api/src/app/nurse"
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// InsertData implements nurse.Repositories
func (repo *repository) InsertData(data nurse.Domain) (string, error) {
	newNurse := MapToNewRecord(data)
	return newNurse.User.ID.String(), repo.DB.Create(&newNurse).Error
}

// SelectAllData implements nurse.Repositories
func (repo *repository) SelectAllData(offset int) ([]nurse.Domain, error) {
	var records []Nurse
	err := repo.DB.Preload("User").Preload("Polyclinic").
		Offset(offset).Limit(10).
		Find(&records).Error
	return MapToBatchDomain(records), err
}

// SelectDataByID implements nurse.Repositories
func (repo *repository) SelectDataByID(id string) (*nurse.Domain, error) {
	var record Nurse
	if err := repo.DB.
		Preload("User").Preload("Polyclinic").
		Where("user_id", id).
		First(&record).Error; err != nil {
		return nil, err
	}

	result := MapToDomain(record)
	return &result, nil
}

// UpdateByID implements nurse.Repositories
func (repo *repository) UpdateByID(id string, domain nurse.Domain) error {
	data := MapToExistingRecord(domain)
	query := repo.DB.Where("user_id", id).Updates(&data)
	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}

	return query.Error
}

// DeleteByID implements nurse.Repositories
func (repo *repository) DeleteByID(id string) error {
	query := repo.DB.Where("user_id", id).Delete(new(Nurse))
	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}

	return query.Error
}

// DeleteUserByID implements nurse.Repositories
func (repo *repository) DeleteUserByID(id string) error {
	query := repo.DB.Where("id", id).Delete(new(User))
	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}

	return query.Error
}

func NewMySQLRepository(DB *gorm.DB) nurse.Repositories {
	return &repository{DB}
}
