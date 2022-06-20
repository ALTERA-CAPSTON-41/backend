package response

import (
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
