package request

import (
	"clinic-api/src/app/prescription"

	"github.com/google/uuid"
)

type NewRequest struct {
	MedicalRecordID string `json:"medical_record_id"`
	Name            string `json:"name"`
	Quantity        int    `json:"quantity"`
	Dosage          string `json:"dosage"`
	Preparatory     string `json:"preparatory"`
	Description     string `json:"description"`
}

type UpdateRequest struct {
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Dosage      string `json:"dosage"`
	Preparatory string `json:"preparatory"`
	Description string `json:"description"`
}

func (req *NewRequest) MapToDomain() prescription.Domain {
	return prescription.Domain{
		MedicalRecordID: uuid.MustParse(req.MedicalRecordID),
		Name:            req.Name,
		Quantity:        req.Quantity,
		Dosage:          req.Dosage,
		Preparatory:     req.Preparatory,
		Description:     req.Description,
	}
}

func (req *UpdateRequest) MapToDomain() prescription.Domain {
	return prescription.Domain{
		Name:        req.Name,
		Quantity:    req.Quantity,
		Dosage:      req.Dosage,
		Preparatory: req.Preparatory,
		Description: req.Description,
	}
}
