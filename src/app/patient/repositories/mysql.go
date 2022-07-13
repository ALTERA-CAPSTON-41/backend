package patient_repositories

import (
	"clinic-api/src/app/patient"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// DeleteByID implements patient.Repositories
func (repo *repository) DeleteByID(id string) error {
	query := repo.DB.Where("id = ?", id).Delete(new(Patient))

	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}
	return query.Error
}

// InsertData implements patient.Repositories
func (repo *repository) InsertData(domain patient.Domain) (id string, err error) {
	record := MapToNewRecord(domain)
	return record.ID.String(), repo.DB.Create(&record).Error
}

// SearchDataByNIKParams implements patient.Repositories
func (repo *repository) SearchDataByNIKParam(nik string) ([]patient.Domain, error) {
	var patients []Patient
	err := repo.DB.Where("nik = ?", nik).Find(&patients).Error
	return MapToBatchDomain(patients), err
}

// SearchDataByNameParams implements patient.Repositories
func (repo *repository) SearchDataByNameParam(name string, offset int) ([]patient.Domain, error) {
	var patients []Patient
	sqlParams := "UPPER(name) LIKE '%" + strings.ToUpper(name) + "%'"
	err := repo.DB.Where(sqlParams).
		Offset(offset).Limit(10).
		Find(&patients).Error
	return MapToBatchDomain(patients), err
}

// SelectAllData implements patient.Repositories
func (repo *repository) SelectAllData(offset int) ([]patient.Domain, error) {
	var patients []Patient
	err := repo.DB.
		Offset(offset).Limit(10).
		Find(&patients).Error
	return MapToBatchDomain(patients), err
}

// SelectDataByID implements patient.Repositories
func (repo *repository) SelectDataByID(id string) (*patient.Domain, error) {
	var record Patient

	if err := repo.DB.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}

	domain := record.MapToDomain()
	return &domain, nil
}

// UpdateByID implements patient.Repositories
func (repo *repository) UpdateByID(id string, domain patient.Domain) error {
	record := MapToExistingRecord(domain)
	query := repo.DB.Where("id = ?", id).Omit("id").Updates(&record)

	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}

	return query.Error
}

func NewMySQLRepository(conn *gorm.DB) patient.Repositories {
	return &repository{conn}
}
