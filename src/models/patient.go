package models

import (
	"clinic-api/src/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenderType string

const (
	MALE GenderType  = "MALE"
	FEMALE GenderType = "FEMALE"
)

type Patient struct {
	ID        uuid.UUID      `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	NIK       string         `json:"nik"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	DOB       time.Time      `json:"dob"`
	Gender    GenderType     `gorm:"type:enum('MALE', 'FEMALE')" json:"gender"`
	BloodType string         `json:"blood_type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PatientRequest struct {
	ID        uuid.UUID      `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	NIK       string         `json:"nik"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	DOB       string         `json:"dob"`
	Gender    GenderType     `gorm:"type:enum('MALE', 'FEMALE')" json:"gender"`
	BloodType string         `json:"blood_type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PatientResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	NIK       string     `json:"nik"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	DOB       string     `json:"dob"`
	Gender    GenderType `json:"gender"`
	BloodType string     `json:"blood_type"`
}

func MapToNewPatient(request PatientRequest) PatientRequest {
	return PatientRequest{
		ID:        uuid.Must(uuid.NewRandom()),
		Name:      request.Name,
		NIK:       request.NIK,
		Phone:     request.Phone,
		Address:   request.Address,
		DOB:       request.DOB,
		Gender:    GenderType(request.Gender),
		BloodType: request.BloodType,
	}
}

func MapToExistingPatient(request PatientRequest) PatientRequest {
	return PatientRequest{
		ID:        request.ID,
		Name:      request.Name,
		NIK:       request.NIK,
		Phone:     request.Phone,
		Address:   request.Address,
		DOB:       request.DOB,
		Gender:    GenderType(request.Gender),
		BloodType: request.BloodType,
	}
}

func MapToPatient(patient Patient) PatientResponse {
	return PatientResponse{
		ID:        patient.ID,
		Name:      patient.Name,
		NIK:       patient.NIK,
		Phone:     patient.Phone,
		Address:   patient.Address,
		DOB:       utils.ConvertDate(patient.DOB),
		Gender:    patient.Gender,
		BloodType: patient.BloodType,
	}
}

func MapToPatientBatch(patients []Patient) []PatientResponse {
	var response []PatientResponse

	for _, patient := range patients {
		response = append(response, MapToPatient(patient))
	}
	return response
}
