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
	CreatedAt        time.Time
	UpdatedAt        time.Time
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
	Age       int
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

type Services interface {
	FindMedicalRecordByID(id string) (*Domain, error)
	FindMedicalRecordByPatientNIK(nik string) ([]Domain, error)
	CreateMedicalRecord(domain Domain) (id string, err error)
}

type Repositories interface {
	SelectPatientIDByNIK(nik string) (string, error)
	SelectDataByPatientID(id string) ([]Domain, error)
	SelectDataByID(id string) (*Domain, error)
	LookupICD10Data(icd10Code string) (ICD10Description string, err error)
	InsertData(domain Domain) (id string, err error)
}
