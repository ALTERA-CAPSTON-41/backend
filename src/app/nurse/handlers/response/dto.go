package response

import (
	"clinic-api/src/app/nurse"
	"clinic-api/src/types"
	"clinic-api/src/utils"

	"github.com/google/uuid"
)

type Response struct {
	ID         uuid.UUID          `json:"id"`
	Name       string             `json:"name"`
	NIP        string             `json:"nip"`
	SIP        string             `json:"sip"`
	Address    string             `json:"address"`
	DOB        string             `json:"dob"`
	Gender     types.GenderEnum   `json:"gender"`
	Email      string             `json:"email"`
	Polyclinic PolyclinicResponse `json:"polyclinic"`
}

type PolyclinicResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Reason interface{} `json:"reason"`
}

func MapToResponse(domain nurse.Domain) Response {
	return Response{
		ID:      domain.User.ID,
		Name:    domain.Name,
		NIP:     domain.NIP,
		SIP:     domain.SIP,
		Address: domain.Address,
		DOB:     utils.ConvertDateToString(domain.DOB),
		Gender:  domain.Gender,
		Email:   domain.User.Email,
		Polyclinic: PolyclinicResponse{
			ID:   domain.Polyclinic.ID,
			Name: domain.Polyclinic.Name,
		},
	}
}

func MapToBatchResponse(domains []nurse.Domain) (responses []Response) {
	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return
}
