package models

import (
	"clinic-api/src/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	ID        uuid.UUID `gorm:"primaryKey;size:191"`
	Name      string
	NIK       string
	Phone     string
	Address   string
	DOB       time.Time
	Gender    GenderType `gorm:"type:enum('MALE', 'FEMALE')"`
	BloodType string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PatientRequest struct {
	Name      string `json:"name"`
	NIK       string `json:"nik"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	DOB       string `json:"dob"`
	Gender    string `json:"gender"`
	BloodType string `json:"blood_type"`
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

func MapToNewPatient(request PatientRequest) Patient {
	return Patient{
		ID:        uuid.Must(uuid.NewRandom()),
		Name:      request.Name,
		NIK:       request.NIK,
		Phone:     request.Phone,
		Address:   request.Address,
		DOB:       utils.ConvertStringToDate(request.DOB),
		Gender:    GenderType(request.Gender),
		BloodType: request.BloodType,
	}
}

func MapToExistingPatient(request PatientRequest, id string) Patient {
	return Patient{
		ID:        uuid.MustParse(id),
		Name:      request.Name,
		NIK:       request.NIK,
		Phone:     request.Phone,
		Address:   request.Address,
		DOB:       utils.ConvertStringToDate(request.DOB),
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
		DOB:       utils.ConvertDateToString(patient.DOB),
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
