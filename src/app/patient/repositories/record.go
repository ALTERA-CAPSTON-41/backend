package patient_repositories

import (
	"clinic-api/src/app/patient"
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;size:191"`
	Name      string
	NIK       string
	Phone     string
	Address   string
	DOB       time.Time
	Gender    types.GenderEnum `gorm:"type:enum('MALE', 'FEMALE')"`
	BloodType string
}

func (record *Patient) MapToDomain() patient.Domain {
	return patient.Domain{
		ID:        record.ID,
		Name:      record.Name,
		NIK:       record.NIK,
		Phone:     record.Phone,
		Address:   record.Address,
		DOB:       record.DOB,
		Gender:    record.Gender,
		BloodType: record.BloodType,
	}
}

func MapToNewRecord(domain patient.Domain) Patient {
	return Patient{
		ID:        uuid.Must(uuid.NewRandom()),
		Name:      domain.Name,
		NIK:       domain.NIK,
		Phone:     domain.Phone,
		Address:   domain.Address,
		DOB:       domain.DOB,
		Gender:    domain.Gender,
		BloodType: domain.BloodType,
	}
}

func MapToExistingRecord(domain patient.Domain) Patient {
	return Patient{
		ID:        uuid.Nil,
		Name:      domain.Name,
		NIK:       domain.NIK,
		Phone:     domain.Phone,
		Address:   domain.Address,
		DOB:       domain.DOB,
		Gender:    domain.Gender,
		BloodType: domain.BloodType,
	}
}

func MapToBatchDomain(records []Patient) []patient.Domain {
	var domains []patient.Domain

	for _, patient := range records {
		domains = append(domains, patient.MapToDomain())
	}
	return domains
}
