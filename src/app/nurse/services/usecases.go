package services

import "clinic-api/src/app/nurse"

type usecase struct {
	repo nurse.Repositories
}

// CreateNurse implements nurse.Services
func (uc *usecase) CreateNurse(nurse nurse.Domain) (string, error) {
	panic("unimplemented")
}

// GetAllNurses implements nurse.Services
func (uc *usecase) GetAllNurses() ([]nurse.Domain, error) {
	panic("unimplemented")
}

// GetNurseByID implements nurse.Services
func (uc *usecase) GetNurseByID(id string) (*nurse.Domain, error) {
	panic("unimplemented")
}

// AmendNurseByID implements nurse.Services
func (uc *usecase) AmendNurseByID(id string, nurse nurse.Domain) error {
	panic("unimplemented")
}

// RemoveNurseByID implements nurse.Services
func (uc *usecase) RemoveNurseByID(id string) error {
	panic("unimplemented")
}

func NewService(repo nurse.Repositories) nurse.Services {
	return &usecase{repo}
}
