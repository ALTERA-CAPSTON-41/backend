package medicalrecord_repositories

import (
	medicalrecord "clinic-api/src/app/medical_record"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecord struct {
	gorm.Model
	ID               uuid.UUID `gorm:"primaryKey;size:191"`
	PatientStatus    types.PatientStatusEnum
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

type ICDResponse struct {
	Description string `json:"Description"`
}

func (mr *MedicalRecord) MapToDomain() medicalrecord.Domain {
	return medicalrecord.Domain{
		ID:               mr.ID,
		PatientStatus:    mr.PatientStatus,
		Symptoms:         mr.Symptoms,
		ICD10Code:        mr.ICD10Code,
		ICD10Description: mr.ICD10Description,
		Suggestions:      mr.Suggestions,
		CreatedAt:        mr.CreatedAt,
		UpdatedAt:        mr.UpdatedAt,
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
		PatientStatus:    domain.PatientStatus,
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
		domain := record.MapToDomain()
		domain.Patient.Age = utils.CountIntervalByYearRoundDown(domain.Patient.DOB, domain.CreatedAt)
		domains = append(domains, domain)
	}
	return domains
}
