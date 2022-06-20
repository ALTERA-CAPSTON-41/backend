package medicalrecord

import (
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID               uuid.UUID
	Symptoms         string
	ICD10Code        string
	ICD10Description string
	Suggestions      string
	Patient          PatientReference
	Doctor           DoctorReference
	Polyclinic       PolyclinicReference
}

type PatientReference struct {
	ID        uuid.UUID
	Name      string
	NIK       string
	Phone     string
	Address   string
	DOB       time.Time
	Gender    types.GenderEnum
	BloodType string
}

type DoctorReference struct {
	ID     uuid.UUID
	Name   string
	NIP    string
	SIP    string
	Gender types.GenderEnum
}

type PolyclinicReference struct {
	ID   int
	Name string
}
