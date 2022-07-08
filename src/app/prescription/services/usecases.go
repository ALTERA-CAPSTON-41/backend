package services

import "clinic-api/src/app/prescription"

type usecase struct {
	repo prescription.Repositories
}

// CreatePrescription implements prescription.Services
func (uc *usecase) CreatePrescription(prescription prescription.Domain) (string, error) {
	return uc.repo.InsertData(prescription)
}

// FindPrescriptionsByID implements prescription.Services
func (uc *usecase) FindPrescriptionsByID(id string) ([]prescription.Domain, error) {
	return uc.repo.SelectAllDataByMedRecordID(id)
}

// AmendPrescriptionByID implements prescription.Services
func (uc *usecase) AmendPrescriptionByID(id string, prescription prescription.Domain) error {
	return uc.repo.UpdateByID(id, prescription)
}

// RemovePrescriptionByID implements prescription.Services
func (uc *usecase) RemovePrescriptionByID(id string) error {
	return uc.repo.DeleteByID(id)
}

func NewService(repo prescription.Repositories) prescription.Services {
	return &usecase{repo}
}
