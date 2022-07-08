package prescription_repositories

import (
	"clinic-api/src/app/prescription"
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// InsertData implements prescription.Repositories
func (repo *repository) InsertData(data prescription.Domain) (string, error) {
	record := MapToNewRecord(data)
	if err := repo.DB.Create(&record).Error; err != nil {
		return "", err
	}

	return record.MedicalRecordID.String(), nil
}

// SelectAllDataByMedRecordID implements prescription.Repositories
func (repo *repository) SelectAllDataByMedRecordID(id string) ([]prescription.Domain, error) {
	var records []Presciption
	if err := repo.DB.Find(&records, "medical_record_id", id).Error; err != nil {
		return nil, err
	}

	return MapToBatchDomain(records), nil
}

// UpdateByID implements prescription.Repositories
func (repo *repository) UpdateByID(id string, domain prescription.Domain) error {
	data := MapToExistingRecord(domain)
	query := repo.DB.Where("id", id).Updates(data)
	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}

	return query.Error
}

// DeleteByID implements prescription.Repositories
func (repo *repository) DeleteByID(id string) error {
	query := repo.DB.Delete(new(Presciption), "id", id)
	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}

	return query.Error
}

func NewMySQLRepository(DB *gorm.DB) prescription.Repositories {
	return &repository{DB}
}
