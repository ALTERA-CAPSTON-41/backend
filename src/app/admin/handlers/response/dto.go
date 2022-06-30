package response

import "clinic-api/src/app/admin"

type Response struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	NIPNumber string `json:"nip"`
	Email     string `json:"email"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Reason string `json:"reason"`
}

func MapToResponse(domain admin.Domain) Response {
	return Response{
		ID:        domain.User.ID.String(),
		Name:      domain.Name,
		NIPNumber: domain.NIP,
		Email:     domain.User.Email,
	}
}

func MapToBatchResponse(domains []admin.Domain) []Response {
	var responses []Response

	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return responses
}
