package services

import "clinic-api/src/app/admin"

type usecase struct {
	repo admin.Repositories
}

// AmendAdminByID implements admin.Services
func (uc *usecase) AmendAdminByID(id string, admin admin.Domain) (err error) {
	return uc.repo.UpdateByID(id, admin)
}

// CreateAdmin implements admin.Services
func (uc *usecase) CreateAdmin(admin admin.Domain) (id string, err error) {
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
	return uc.repo.DeleteByID(id)
}

func NewService(repo admin.Repositories) admin.Services {
	return &usecase{repo}
}
