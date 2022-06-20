package medicalrecord_repositories

import (
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
)

type MedicalRecord struct {
	ID               uuid.UUID `gorm:"primaryKey;size:191"`
	Symptoms         string
	ICD10Code        string
	ICD10Description string
	Suggestions      string
	PolyclinicID     int
	PatientID        uuid.UUID `gorm:"size:191"`
	DoctorID         uuid.UUID `gorm:"size:191"`
	Polyclinic       Polyclinic
	Patient          Patient
	Doctor           Doctor `gorm:"references:UserID"`
}

type Polyclinic struct {
	ID   int
	Name string
}

type Patient struct {
	ID        uuid.UUID `gorm:"primaryKey;size:191"`
	Name      string
	NIK       string
	Phone     string
	Address   string
	DOB       time.Time
	Gender    types.GenderEnum `gorm:"type:enum('MALE', 'FEMALE')"`
	BloodType string
}

type Doctor struct {
	UserID uuid.UUID `gorm:"primaryKey;size:191"`
	Name   string
	NIP    string           `gorm:"column:nip"`
	SIP    string           `gorm:"column:sip"`
	Gender types.GenderEnum `gorm:"type:enum('MALE', 'FEMALE')"`
}
