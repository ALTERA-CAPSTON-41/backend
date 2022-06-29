package services

import (
	"clinic-api/src/app/admin"
	"errors"
)

type usecase struct {
	repo admin.Repositories
}

// AmendAdminByID implements admin.Services
func (uc *usecase) AmendAdminByID(id string, admin admin.Domain) (err error) {
	return uc.repo.UpdateByID(id, admin)
}

// CreateAdmin implements admin.Services
func (uc *usecase) CreateAdmin(admin admin.Domain) (id string, err error) {
	email, err := uc.repo.LookupDataByEmail(admin.User.Email)
	if err != nil {
		return "", err
	}

	if email != "" {
		return "", errors.New("email is already used")
	}

	return uc.repo.InsertData(admin)
}

// GetAdminByID implements admin.Services
func (uc *usecase) GetAdminByID(id string) (admin *admin.Domain, err error) {
	return uc.repo.SelectDataByID(id)
}

// GetAllAdmins implements admin.Services
func (uc *usecase) GetAllAdmins() (admins []admin.Domain, err error) {
	return uc.repo.SelectAllData()
}

// RemoveAdminByID implements admin.Services
func (uc *usecase) RemoveAdminByID(id string) (err error) {
	if err := uc.repo.DeleteByID(id); err != nil {
		return err
	}
	return uc.repo.DeleteUserByID(id)
}

func NewService(repo admin.Repositories) admin.Services {
	return &usecase{repo}
}
