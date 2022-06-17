package request

import (
	"clinic-api/src/app/nurse"
	"clinic-api/src/types"
	"clinic-api/src/utils"
)

type NewRequest struct {
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

type UpdateRequest struct {
	Name         string `json:"name"`
	NIP          string `json:"nip"`
	SIP          string `json:"sip"`
	Address      string `json:"address"`
	DOB          string `json:"dob"`
	Gender       string `json:"gender"`
	PolyclinicID int    `json:"polyclinic_id"`
}

func (req *NewRequest) MapToDomain() nurse.Domain {
	return nurse.Domain{
		Name:    req.Name,
		NIP:     req.NIP,
		SIP:     req.SIP,
		Address: req.Address,
		DOB:     utils.ConvertStringToDate(req.DOB),
		Gender:  types.GenderEnum(req.Gender),
		Polyclinic: nurse.PolyclinicReference{
			ID: req.PolyclinicID,
		},
		User: nurse.UserReference{
			Email:    req.Email,
			Password: req.Password,
		},
	}
}

func (req *UpdateRequest) MapToDomain() nurse.Domain {
	return nurse.Domain{
		Name:    req.Name,
		NIP:     req.NIP,
		SIP:     req.SIP,
		Address: req.Address,
		DOB:     utils.ConvertStringToDate(req.DOB),
		Gender:  types.GenderEnum(req.Gender),
		Polyclinic: nurse.PolyclinicReference{
			ID: req.PolyclinicID,
		},
	}
}
