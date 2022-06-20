package response

import (
	medicalrecord "clinic-api/src/app/medical_record"
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
)

type Response struct {
	ID               uuid.UUID           `json:"id"`
	Symptoms         string              `json:"symptoms"`
	ICD10Code        string              `json:"icd10_code"`
	ICD10Description string              `json:"icd10_description"`
	Suggestions      string              `json:"suggestions"`
	Patient          PatientReference    `json:"patient"`
	Doctor           DoctorReference     `json:"doctor"`
	Polyclinic       PolyclinicReference `json:"polyclinic"`
}

type PatientReference struct {
	ID        uuid.UUID        `json:"id"`
	Name      string           `json:"name"`
	NIK       string           `json:"nik"`
	Phone     string           `json:"phone"`
	Address   string           `json:"address"`
	DOB       time.Time        `json:"dob"`
	Gender    types.GenderEnum `json:"gender"`
	BloodType string           `json:"blood_type"`
}

type DoctorReference struct {
	ID     uuid.UUID        `json:"id"`
	Name   string           `json:"name"`
	NIP    string           `json:"nip"`
	SIP    string           `json:"sip"`
	Gender types.GenderEnum `json:"gender"`
}

type PolyclinicReference struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateResponse struct {
	ID uuid.UUID `json:"id"`
}

func MapToResponse(domain medicalrecord.Domain) Response {
	return Response{
		ID:               domain.ID,
		Symptoms:         domain.Symptoms,
		ICD10Code:        domain.ICD10Code,
		ICD10Description: domain.ICD10Description,
		Suggestions:      domain.Suggestions,
		Patient: PatientReference{
			ID:        domain.Patient.ID,
			Name:      domain.Patient.Name,
			NIK:       domain.Patient.NIK,
			Phone:     domain.Patient.Phone,
			Address:   domain.Patient.Address,
			DOB:       domain.Patient.DOB,
			Gender:    domain.Patient.Gender,
			BloodType: domain.Patient.BloodType,
		},
		Doctor: DoctorReference{
			ID:     domain.Doctor.ID,
			Name:   domain.Doctor.Name,
			NIP:    domain.Doctor.NIP,
			SIP:    domain.Doctor.SIP,
			Gender: domain.Doctor.Gender,
		},
		Polyclinic: PolyclinicReference{
			ID:   domain.Polyclinic.ID,
			Name: domain.Polyclinic.Name,
		},
	}
}

func MapToBatchResponse(domains []medicalrecord.Domain) []Response {
	var responses []Response

	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return responses
}
