package response

import (
	"clinic-api/src/app/patient"
	"clinic-api/src/types"
	"clinic-api/src/utils"

	"github.com/google/uuid"
)

type Response struct {
	ID        uuid.UUID        `json:"id"`
	Name      string           `json:"name"`
	NIK       string           `json:"nik"`
	Phone     string           `json:"phone"`
	Address   string           `json:"address"`
	DOB       string           `json:"dob"`
	Gender    types.GenderEnum `  json:"gender"`
	BloodType string           `json:"blood_type"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Reason string `json:"reason"`
}

func MapToResponse(domain patient.Domain) Response {
	return Response{
		ID:        domain.ID,
		Name:      domain.Name,
		NIK:       domain.NIK,
		Phone:     domain.Phone,
		Address:   domain.Address,
		DOB:       utils.ConvertDateToString(domain.DOB),
		Gender:    domain.Gender,
		BloodType: domain.BloodType,
	}
}

func MapToBatchResponse(domains []patient.Domain) []Response {
	var responses []Response

	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return responses
}
