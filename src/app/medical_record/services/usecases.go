package services

import medicalrecord "clinic-api/src/app/medical_record"

type usecase struct {
	repo medicalrecord.Repositories
}

// CreateMedicalRecord implements medicalrecord.Services
func (uc *usecase) CreateMedicalRecord(domain medicalrecord.Domain) (id string, err error) {
	panic("unimplemented")
}

// FindMedicalRecordByID implements medicalrecord.Services
func (uc *usecase) FindMedicalRecordByID(id string) (*medicalrecord.Domain, error) {
	panic("unimplemented")
}

// FindMedicalRecordByPatientNIK implements medicalrecord.Services
func (uc *usecase) FindMedicalRecordByPatientNIK(nik string) ([]medicalrecord.Domain, error) {
	panic("unimplemented")
}

func NewServices(repo medicalrecord.Repositories) medicalrecord.Services {
	return &usecase{repo}
}
