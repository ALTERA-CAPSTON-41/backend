package prescription_repositories

import (
	"clinic-api/src/app/prescription"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Presciption struct {
	gorm.Model
	ID              uuid.UUID `gorm:"size:191"`
	MedicalRecordID uuid.UUID `gorm:"size:191"`
	Name            string
	Quantity        int
	Dosage          string
	Preparatory     string
	Description     string
	MedicalRecord   MedicalRecord
}

type MedicalRecord struct {
	ID uuid.UUID
}

func MapToDomain(record Presciption) prescription.Domain {
	return prescription.Domain{
		ID:          record.ID,
		Name:        record.Name,
		Quantity:    record.Quantity,
		Dosage:      record.Dosage,
		Preparatory: record.Preparatory,
		Description: record.Description,
	}
}

func MapToNewRecord(domain prescription.Domain) Presciption {
	return Presciption{
		ID:              uuid.Must(uuid.NewRandom()),
		MedicalRecordID: domain.MedicalRecordID,
		Name:            domain.Name,
		Quantity:        domain.Quantity,
		Dosage:          domain.Dosage,
		Preparatory:     domain.Preparatory,
		Description:     domain.Description,
	}
}

func MapToExistingRecord(domain prescription.Domain) Presciption {
	return Presciption{
		Name:        domain.Name,
		Quantity:    domain.Quantity,
		Dosage:      domain.Dosage,
		Preparatory: domain.Preparatory,
		Description: domain.Description,
	}
}

func MapToBatchDomain(records []Presciption) (domains []prescription.Domain) {
	for _, record := range records {
		domains = append(domains, MapToDomain(record))
	}
	return
}
