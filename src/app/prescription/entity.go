package prescription

import "github.com/google/uuid"

type Domain struct {
	ID              uuid.UUID
	MedicalRecordID uuid.UUID
	Name            string
	Quantity        int
	Dosage          string
	Preparatory     string
	Description     string
}

type Services interface {
	CreatePrescription(prescription Domain) (string, error)
	FindPrescriptionsByID(id string) ([]Domain, error)
	AmendPrescriptionByID(id string, prescription Domain) error
	RemovePrescriptionByID(id string) error
}

type Repositories interface {
	InsertData(data Domain) (string, error)
	SelectAllDataByMedRecordID(id string) ([]Domain, error)
	UpdateByID(id string, domain Domain) error
	DeleteByID(id string) error
}
