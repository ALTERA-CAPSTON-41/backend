package services

import (
	medicalrecord "clinic-api/src/app/medical_record"
)

type usecase struct {
	repo medicalrecord.Repositories
}

// CreateMedicalRecord implements medicalrecord.Services
func (uc *usecase) CreateMedicalRecord(domain medicalrecord.Domain) (id string, err error) {
	if domain.ICD10Description, err = uc.repo.LookupICD10Data(domain.ICD10Code); err != nil {
		return "", err
	}
	return uc.repo.InsertData(domain)
}

// FindMedicalRecordByID implements medicalrecord.Services
func (uc *usecase) FindMedicalRecordByID(id string) (*medicalrecord.Domain, error) {
	return uc.repo.SelectDataByID(id)
}

// FindMedicalRecordByPatientNIK implements medicalrecord.Services
func (uc *usecase) FindMedicalRecordByPatientNIK(nik string) ([]medicalrecord.Domain, error) {
	id, err := uc.repo.SelectPatientIDByNIK(nik)
	if err != nil {
		return nil, err
	}

	return uc.repo.SelectDataByPatientID(id)
}

func NewServices(repo medicalrecord.Repositories) medicalrecord.Services {
	return &usecase{repo}
}
