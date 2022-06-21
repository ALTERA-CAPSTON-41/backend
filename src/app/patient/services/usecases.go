package services

import "clinic-api/src/app/patient"

type usecase struct {
	repo patient.Repositories
}

// AmendPatientByID implements patient.Services
func (uc *usecase) AmendPatientByID(id string, domain patient.Domain) error {
	return uc.repo.UpdateByID(id, domain)
}

// CreatePatient implements patient.Services
func (uc *usecase) CreatePatient(domain patient.Domain) (id string, err error) {
	return uc.repo.InsertData(domain)
}

// GetPatientByID implements patient.Services
func (uc *usecase) GetPatientByID(id string) (*patient.Domain, error) {
	return uc.repo.SelectDataByID(id)
}

// HuntPatientByNameOrNIKOrAll implements patient.Services
func (uc *usecase) HuntPatientByNameOrNIKOrAll(domain patient.Domain, page int) ([]patient.Domain, error) {
	offset := (page - 1) * 10
	if domain.NIK != "" {
		return uc.repo.SearchDataByNIKParam(domain.NIK)
	}

	if domain.Name != "" {
		return uc.repo.SearchDataByNameParam(domain.Name, offset)
	}

	return uc.repo.SelectAllData(offset)
}

// RemovePatientByID implements patient.Services
func (uc *usecase) RemovePatientByID(id string) error {
	return uc.repo.DeleteByID(id)
}

func NewService(repo patient.Repositories) patient.Services {
	return &usecase{repo}
}
