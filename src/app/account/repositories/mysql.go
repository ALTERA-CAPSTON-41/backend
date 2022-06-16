package account_repositories

import (
	"clinic-api/src/app/account"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// LookupAccountByEmail implements account.Repositories
func (repo *repository) LookupAccountByEmail(email string) (*account.Domain, error) {
	var record User

	if err := repo.DB.Where("email = ?", email).Find(&record).Error; err != nil {
		return nil, err
	}

	if record.ID == uuid.Nil {
		return nil, errors.New("record not found")
	}

	result := record.MapToDomain()
	return &result, nil
}

// LookupAdminByUserID implements account.Repositories
func (repo *repository) LookupAdminByUserID(id string) (*account.UserDataDomain, error) {
	var record Admin

	if err := repo.DB.Where("user_id = ?", id).Find(&record).Error; err != nil {
		return nil, err
	}

	result := record.MapToDomain()
	return &result, nil
}

// LookupDoctorByUserID implements account.Repositories
func (repo *repository) LookupDoctorByUserID(id string) (*account.UserDataDomain, error) {
	var record Doctor

	if err := repo.DB.Where("user_id = ?", id).Find(&record).Error; err != nil {
		return nil, err
	}

	result := record.MapToDomain()
	return &result, nil
}

func NewMySQLRepository(DB *gorm.DB) account.Repositories {
	return &repository{DB}
}
