package request

import (
	"clinic-api/src/app/patient"
	"clinic-api/src/types"
	"clinic-api/src/utils"
)

type Request struct {
	Name      string `json:"name"`
	NIK       string `json:"nik" validate:"numeric,min=16"`
	Phone     string `json:"phone" validate:"numeric,min=10"`
	Address   string `json:"address"`
	DOB       string `json:"dob"`
	Gender    string `json:"gender"`
	BloodType string `json:"blood_type"`
}

func (req *Request) MapToDomain() patient.Domain {
	return patient.Domain{
		Name:      req.Name,
		NIK:       req.NIK,
		Phone:     req.Phone,
		Address:   req.Address,
		DOB:       utils.ConvertStringToDate(req.DOB),
		Gender:    types.GenderEnum(req.Gender),
		BloodType: req.BloodType,
	}
}
