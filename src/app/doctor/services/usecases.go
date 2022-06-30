package services

import (
	"clinic-api/src/app/doctor"
	"errors"
)

type usecase struct {
	repo doctor.Repositories
}

// AmendDoctorByID implements doctor.Services
func (uc *usecase) AmendDoctorByID(id string, doctor doctor.Domain) error {
	return uc.repo.UpdateByID(id, doctor)
}

// CreateDoctor implements doctor.Services
func (uc *usecase) CreateDoctor(doctor doctor.Domain) (string, error) {
	email, err := uc.repo.LookupDataByEmail(doctor.User.Email)
	if err != nil {
		return "", err
	}

	if email != "" {
		return "", errors.New("email is already used")
	}

	return uc.repo.InsertData(doctor)
}

// GetAllDoctors implements doctor.Services
func (uc *usecase) GetAllDoctors(page int) ([]doctor.Domain, error) {
	offset := (page - 1) * 10
	return uc.repo.SelectAllData(offset)
}

// GetDoctorByID implements doctor.Services
func (uc *usecase) GetDoctorByID(id string) (*doctor.Domain, error) {
	return uc.repo.SelectDataByID(id)
}

// RemoveDoctorByID implements doctor.Services
func (uc *usecase) RemoveDoctorByID(id string) error {
	if err := uc.repo.DeleteByID(id); err != nil {
		return err
	}
	return uc.repo.DeleteUserByID(id)
}

func NewService(repo doctor.Repositories) doctor.Services {
	return &usecase{repo}
}
