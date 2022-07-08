package response

import (
	"clinic-api/src/app/prescription"

	"github.com/google/uuid"
)

type Response struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	Dosage      string    `json:"dosage"`
	Preparatory string    `json:"preparatory"`
	Description string    `json:"description"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

func MapToResponse(domain prescription.Domain) Response {
	return Response{
		ID:          domain.ID,
		Name:        domain.Name,
		Quantity:    domain.Quantity,
		Dosage:      domain.Dosage,
		Preparatory: domain.Preparatory,
		Description: domain.Description,
	}
}

func MapToBatchResponse(domains []prescription.Domain) (responses []Response) {
	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return
}
