package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type genderType string

const (
	MALE   = "MALE"
	FEMALE = "FEMALE"
)

type Patient struct {
	ID        uuid.UUID      `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	NIK       string         `json:"nik"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	DOB       time.Time      `json:"dob"`
	Gender    genderType     `gorm:"type:enum('MALE', 'FEMALE')" json:"gender"`
	BloodType string         `json:"blood_type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PatientResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	NIK       string    `json:"nik"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	DOB       time.Time `json:"dob"`
	Gender    genderType    `json:"gender"`
	BloodType string    `json:"blood_type"`
}

func MapToNewPatient(model Patient) Patient {
	return Patient{
		ID:        uuid.Must(uuid.NewRandom()),
		Name:      model.Name,
		NIK:       model.NIK,
		Phone:     model.Phone,
		Address:   model.Address,
		DOB:       model.DOB,
		Gender:    model.Gender,
		BloodType: model.BloodType,
	}
}

func MapToExistingPatient(model Patient) Patient {
	return Patient{
		ID:        model.ID,
		Name:      model.Name,
		NIK:       model.NIK,
		Phone:     model.Phone,
		Address:   model.Address,
		DOB:       model.DOB,
		Gender:    model.Gender,
		BloodType: model.BloodType,
	}
}

func MapToPatient(model Patient) PatientResponse {
	return PatientResponse{
		ID:        model.ID,
		Name:      model.Name,
		NIK:       model.NIK,
		Phone:     model.Phone,
		Address:   model.Address,
		DOB:       model.DOB,
		Gender:    model.Gender,
		BloodType: model.BloodType,
	}
}

func MapToPatientBatch(model []Patient) []PatientResponse {
	var patients []PatientResponse

	for _, patient := range model {
		patients = append(patients, MapToPatient(patient))
	}
	return patients
}
