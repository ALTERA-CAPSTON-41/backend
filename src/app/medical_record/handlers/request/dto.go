package request

import (
	medicalrecord "clinic-api/src/app/medical_record"
	"clinic-api/src/types"

	"github.com/google/uuid"
)

type Request struct {
	PatientStatus string `json:"patient_status"`
	Symptoms      string `json:"symptoms"`
	ICD10Code     string `json:"icd10_code"`
	Suggestions   string `json:"suggestions"`
	PatientID     string `json:"patient_id"`
	PolyclinicID  int    `json:"polyclinic_id"`
}

type UpdateRequest struct {
	PatientStatus string `json:"patient_status"`
	Symptoms      string `json:"symptoms"`
	ICD10Code     string `json:"icd10_code"`
	Suggestions   string `json:"suggestions"`
}

func (req *Request) MapToDomain() medicalrecord.Domain {
	return medicalrecord.Domain{
		PatientStatus: types.PatientStatusEnum(req.PatientStatus),
		Symptoms:      req.Symptoms,
		ICD10Code:     req.ICD10Code,
		Suggestions:   req.Suggestions,
		Patient: medicalrecord.PatientReference{
			ID: uuid.MustParse(req.PatientID),
		},
		Polyclinic: medicalrecord.PolyclinicReference{
			ID: req.PolyclinicID,
		},
	}
}

func (req *UpdateRequest) MapToDomain() medicalrecord.Domain {
	return medicalrecord.Domain{
		PatientStatus: types.PatientStatusEnum(req.PatientStatus),
		Symptoms:      req.Symptoms,
		ICD10Code:     req.ICD10Code,
		Suggestions:   req.Suggestions,
	}
}
