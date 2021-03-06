package admin_repositories

import (
	"clinic-api/src/app/admin"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// DeleteByID implements admin.Repositories
func (repo *repository) DeleteByID(id string) (err error) {
	deletion := repo.DB.Where("user_id = ?", id).Delete(new(Admin))

	if deletion.RowsAffected == 0 && deletion.Error == nil {
		return errors.New("record not found")
	}
	return deletion.Error
}

// DeleteUserByID implements admin.Repositories
func (repo *repository) DeleteUserByID(id string) (err error) {
	deletion := repo.DB.Where("id", id).Delete(new(User))

	if deletion.RowsAffected == 0 && deletion.Error == nil {
		return errors.New("record not found")
	}
	return deletion.Error
}

// InsertData implements admin.Repositories
func (repo *repository) InsertData(data admin.Domain) (id string, err error) {
	record := MapToNewRecord(data)

	if err = repo.DB.Create(&record).Error; err != nil {
		return uuid.Nil.String(), err
	}
	return record.UserID.String(), nil
}

// SelectAllData implements admin.Repositories
func (repo *repository) SelectAllData() (data []admin.Domain, err error) {
	var records []Admin

	if err = repo.DB.Preload("User").Find(&records).Error; err != nil {
		return nil, err
	}
	return MapToBatchDomain(records), nil
}

// SelectDataByID implements admin.Repositories
func (repo *repository) SelectDataByID(id string) (data *admin.Domain, err error) {
	var record Admin

	if err = repo.DB.Preload("User").Where("user_ID = ?", id).First(&record).Error; err != nil {
		return nil, err
	}
	admin := MapToDomain(record)
	return &admin, nil
}

// LookupDataByEmail implements admin.Repositories
func (repo *repository) LookupDataByEmail(email string) (string, error) {
	var result string
	if err := repo.DB.Table("users").
		Select("email").Where("email = ?", email).
		Find(&result).Error; err != nil {
		return "", err
	}
	return result, nil
}

// UpdateByID implements admin.Repositories
func (repo *repository) UpdateByID(id string, data admin.Domain) (err error) {
	record := MapToExistingRecord(data)
	alteration := repo.DB.Model(new(Admin)).Where("user_ID = ?", id).Updates(&record)

	if alteration.RowsAffected == 0 && alteration.Error == nil {
		return errors.New("record not found")
	}
	return alteration.Error
}

func NewMySQLRepository(DB *gorm.DB) admin.Repositories {
	return &repository{DB}
}
