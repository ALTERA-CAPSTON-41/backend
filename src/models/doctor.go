package models

import (
	"clinic-api/src/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Doctor struct {
	UserID       uuid.UUID `gorm:"primaryKey;size:191"`
	Name         string
	NIP          string `gorm:"column:nip"`
	SIP          string `gorm:"column:sip"`
	Address      string
	DOB          time.Time
	Gender       GenderType `gorm:"type:enum('MALE', 'FEMALE')"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	PolyclinicID int
	Polyclinic   Polyclinic
	User         User
}

type DoctorRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	NIP          string `json:"nip"`
	SIP          string `json:"sip"`
	Address      string `json:"address"`
	DOB          string `json:"dob"`
	Gender       string `json:"gender"`
	PolyclinicID int    `json:"polyclinic_id"`
}

type DoctorResponse struct {
	ID         uuid.UUID          `json:"id"`
	Name       string             `json:"name"`
	NIP        string             `json:"nip"`
	SIP        string             `json:"sip"`
	Address    string             `json:"address"`
	DOB        string             `json:"dob"`
	Gender     GenderType         `json:"gender"`
	Email      string             `json:"email"`
	Polyclinic PolyclinicResponse `json:"polyclinic"`
}

func MapToNewDoctor(request DoctorRequest) Doctor {
	password, _ := utils.CreateHash(request.Password)
	userID := uuid.Must(uuid.NewRandom())
	return Doctor{
		User:         User{ID: userID, Email: request.Email, Password: password, Role: DOCTOR},
		UserID:       userID,
		Name:         request.Name,
		NIP:          request.NIP,
		SIP:          request.SIP,
		Address:      request.Address,
		DOB:          utils.ConvertStringToDate(request.DOB),
		Gender:       MALE,
		PolyclinicID: request.PolyclinicID,
	}
}

func MapToExistingDoctor(request DoctorRequest, id string) Doctor {
	return Doctor{
		UserID:       uuid.MustParse(id),
		Name:         request.Name,
		NIP:          request.NIP,
		SIP:          request.SIP,
		Address:      request.Address,
		DOB:          utils.ConvertStringToDate(request.DOB),
		Gender:       GenderType(request.Gender),
		PolyclinicID: request.PolyclinicID,
	}
}

func MapToDoctorResponse(doctor Doctor) DoctorResponse {
	return DoctorResponse{
		ID:      doctor.UserID,
		Name:    doctor.Name,
		NIP:     doctor.NIP,
		SIP:     doctor.SIP,
		Address: doctor.Address,
		DOB:     utils.ConvertDateToString(doctor.DOB),
		Gender:  doctor.Gender,
		Email:   doctor.User.Email,
		Polyclinic: PolyclinicResponse{
			doctor.PolyclinicID, doctor.Polyclinic.Name,
		},
	}
}

func MapToDoctorBatchResponse(doctors []Doctor) []DoctorResponse {
	var response []DoctorResponse

	for _, doctor := range doctors {
		response = append(response, MapToDoctorResponse(doctor))
	}
	return response
}
