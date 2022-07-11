package services

import (
	"clinic-api/src/app/nurse"
	"errors"
)

type usecase struct {
	repo nurse.Repositories
}

// CreateNurse implements nurse.Services
func (uc *usecase) CreateNurse(nurse nurse.Domain) (string, error) {
	email, err := uc.repo.LookupDataByEmail(nurse.User.Email)
	if err != nil {
		return "", err
	}

	if email != "" {
		return "", errors.New("email is already used")
	}

	return uc.repo.InsertData(nurse)
}

// GetAllNurses implements nurse.Services
func (uc *usecase) GetAllNurses(polyclinic, page int) ([]nurse.Domain, error) {
	offset := (page - 1) * 10
	return uc.repo.SelectAllData(polyclinic, offset)
}

// GetNurseByID implements nurse.Services
func (uc *usecase) GetNurseByID(id string) (*nurse.Domain, error) {
	return uc.repo.SelectDataByID(id)
}

// AmendNurseByID implements nurse.Services
func (uc *usecase) AmendNurseByID(id string, nurse nurse.Domain) error {
	return uc.repo.UpdateByID(id, nurse)
}

// RemoveNurseByID implements nurse.Services
func (uc *usecase) RemoveNurseByID(id string) error {
	if err := uc.repo.DeleteByID(id); err != nil {
		return err
	}

	return uc.repo.DeleteUserByID(id)
}

func NewService(repo nurse.Repositories) nurse.Services {
	return &usecase{repo}
}
