package medicalrecord_repositories

import (
	medicalrecord "clinic-api/src/app/medical_record"
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

func (mr *MedicalRecord) MapToDomain() medicalrecord.Domain {
	return medicalrecord.Domain{
		ID:               mr.ID,
		Symptoms:         mr.Symptoms,
		ICD10Code:        mr.ICD10Code,
		ICD10Description: mr.ICD10Description,
		Suggestions:      mr.Suggestions,
		Patient: medicalrecord.PatientReference{
			ID:        mr.Patient.ID,
			Name:      mr.Patient.Name,
			NIK:       mr.Patient.NIK,
			Phone:     mr.Patient.Phone,
			Address:   mr.Patient.Address,
			DOB:       mr.Patient.DOB,
			Gender:    mr.Patient.Gender,
			BloodType: mr.Patient.BloodType,
		},
		Doctor: medicalrecord.DoctorReference{
			ID:     mr.Doctor.UserID,
			Name:   mr.Doctor.Name,
			NIP:    mr.Doctor.NIP,
			SIP:    mr.Doctor.SIP,
			Gender: mr.Doctor.Gender,
		},
		Polyclinic: medicalrecord.PolyclinicReference{
			ID:   mr.Polyclinic.ID,
			Name: mr.Polyclinic.Name,
		},
	}
}

func MapToNewRecord(domain medicalrecord.Domain) MedicalRecord {
	return MedicalRecord{
		ID:               uuid.Must(uuid.NewRandom()),
		Symptoms:         domain.Symptoms,
		ICD10Code:        domain.ICD10Code,
		ICD10Description: domain.ICD10Description,
		Suggestions:      domain.Suggestions,
		PolyclinicID:     domain.Polyclinic.ID,
		PatientID:        domain.Patient.ID,
		DoctorID:         domain.Doctor.ID,
	}
}

func MapToBatchDomain(records []MedicalRecord) []medicalrecord.Domain {
	var domains []medicalrecord.Domain

	for _, record := range records {
		domains = append(domains, record.MapToDomain())
	}
	return domains
}
