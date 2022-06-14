package services

import (
	"clinic-api/src/app/doctor"
)

type usecase struct {
	repo doctor.Repositories
}

// AmendDoctorByID implements doctor.Services
func (uc *usecase) AmendDoctorByID(id string, doctor doctor.Domain) error {
	panic("unimplemented")
}

// CreateDoctor implements doctor.Services
func (uc *usecase) CreateDoctor(doctor doctor.Domain) (string, error) {
	panic("unimplemented")
}

// GetAllDoctors implements doctor.Services
func (uc *usecase) GetAllDoctors() ([]doctor.Domain, error) {
	panic("unimplemented")
}

// GetDoctorByID implements doctor.Services
func (uc *usecase) GetDoctorByID(id string) (*doctor.Domain, error) {
	panic("unimplemented")
}

// RemoveDoctorByID implements doctor.Services
func (uc *usecase) RemoveDoctorByID(id string) error {
	panic("unimplemented")
}

func NewService(repo doctor.Repositories) doctor.Services {
	return &usecase{repo}
}
