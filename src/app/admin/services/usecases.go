package services

import "clinic-api/src/app/admin"

type usecase struct {
	repo admin.Repositories
}

// AmendAdminByID implements admin.Services
func (uc *usecase) AmendAdminByID(id string, admin admin.Domain) (err error) {
	panic("unimplemented")
}

// CreateAdmin implements admin.Services
func (uc *usecase) CreateAdmin(admin admin.Domain) (id string, err error) {
	panic("unimplemented")
}

// GetAdminByID implements admin.Services
func (uc *usecase) GetAdminByID(id string) (admin admin.Domain, err error) {
	panic("unimplemented")
}

// GetAllAdmins implements admin.Services
func (uc *usecase) GetAllAdmins() (admins []admin.Domain, err error) {
	panic("unimplemented")
}

// RemoveAdminByID implements admin.Services
func (uc *usecase) RemoveAdminByID(id string) (err error) {
	panic("unimplemented")
}

func NewService(repo admin.Repositories) admin.Services {
	return &usecase{repo}
}
