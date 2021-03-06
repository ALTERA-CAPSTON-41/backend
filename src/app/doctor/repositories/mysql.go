package doctor_repositories

import (
	"clinic-api/src/app/doctor"
	"errors"
	"fmt"

	"github.com/google/uuid"
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

// DeleteUserByID implements doctor.Repositories
func (repo *repository) DeleteUserByID(id string) error {
	query := repo.DB.Where("id", id).Delete(new(User))

	if query.RowsAffected < 1 && query.Error == nil {
		return errors.New("record not found")
	}
	return query.Error
}

// LookupDataByEmail implements doctor.Repositories
func (repo *repository) LookupDataByEmail(email string) (string, error) {
	var result string
	if err := repo.DB.Table("users").
		Select("email").Where("email = ?", email).Find(&result).Error; err != nil {
		return "", err
	}
	return result, nil
}

// InsertData implements doctor.Repositories
func (repo *repository) InsertData(data doctor.Domain) (string, error) {
	newDoctor := MapToNewRecord(data)
	return newDoctor.User.ID.String(), repo.DB.Create(&newDoctor).Error
}

// SelectAllData implements doctor.Repositories
func (repo *repository) SelectAllData(polyclinic, offset int) ([]doctor.Domain, error) {
	var (
		records []Doctor
		query   string
	)
	if polyclinic != 0 {
		query = fmt.Sprintf("polyclinic_id = %d", polyclinic)
	}
	err := repo.DB.Preload("User").Preload("Polyclinic").
		Offset(offset).Limit(10).
		Find(&records, query).Error
	return MapToBatchDomain(records), err
}

// SelectDataByID implements doctor.Repositories
func (repo *repository) SelectDataByID(id string) (*doctor.Domain, error) {
	var record Doctor
	err := repo.DB.Preload("User").Preload("Polyclinic").Where("user_id", id).Find(&record).Error

	if err == nil && record.User.ID == uuid.Nil {
		return nil, errors.New("record not found")
	}
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
