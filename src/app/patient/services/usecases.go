package services

import "clinic-api/src/app/patient"

type usecase struct {
	repo patient.Repositories
}

// AmendPatientByID implements patient.Services
func (uc *usecase) AmendPatientByID(id string, domain patient.Domain) error {
	panic("unimplemented")
}

// CreatePatient implements patient.Services
func (uc *usecase) CreatePatient(domain patient.Domain) (id string, err error) {
	panic("unimplemented")
}

// GetPatientByID implements patient.Services
func (uc *usecase) GetPatientByID(id string) (*patient.Domain, error) {
	panic("unimplemented")
}

// HuntPatientByNameOrNIKOrAll implements patient.Services
func (uc *usecase) HuntPatientByNameOrNIKOrAll(domain patient.Domain) ([]patient.Domain, error) {
	panic("unimplemented")
}

// RemovePatientByID implements patient.Services
func (uc *usecase) RemovePatientByID(id string) error {
	panic("unimplemented")
}

func NewService(repo patient.Repositories) patient.Services {
	return &usecase{repo}
}
