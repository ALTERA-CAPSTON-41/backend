package request

import (
	"clinic-api/src/app/doctor"
	"clinic-api/src/types"
	"clinic-api/src/utils"
)

type NewRequest struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	Name         string `json:"name"`
	NIP          string `json:"nip"`
	SIP          string `json:"sip"`
	Address      string `json:"address"`
	DOB          string `json:"dob"`
	Gender       string `json:"gender"`
	PolyclinicID int    `json:"polyclinic_id"`
}

type UpdateRequest struct {
	Name         string `json:"name"`
	NIP          string `json:"nip"`
	SIP          string `json:"sip"`
	Address      string `json:"address"`
	DOB          string `json:"dob"`
	Gender       string `json:"gender"`
	PolyclinicID int    `json:"polyclinic_id"`
}

func (req *NewRequest) MapToDomain() doctor.Domain {
	return doctor.Domain{
		Name:    req.Name,
		NIP:     req.NIP,
		SIP:     req.SIP,
		Address: req.Address,
		DOB:     utils.ConvertStringToDate(req.DOB),
		Gender:  types.GenderEnum(req.Gender),
		Polyclinic: doctor.PolyclinicReference{
			ID: req.PolyclinicID,
		},
		User: doctor.UserReference{
			Email:    req.Email,
			Password: req.Password,
		},
	}
}

func (req *UpdateRequest) MapToDomain() doctor.Domain {
	return doctor.Domain{
		Name:    req.Name,
		NIP:     req.NIP,
		SIP:     req.SIP,
		Address: req.Address,
		DOB:     utils.ConvertStringToDate(req.DOB),
		Gender:  types.GenderEnum(req.Gender),
		Polyclinic: doctor.PolyclinicReference{
			ID: req.PolyclinicID,
		},
	}
}
