package response

import "clinic-api/src/app/admin"

type Response struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	NIPNumber string `json:"nip"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

func MapToResponse(domain admin.Domain) Response {
	return Response{
		ID:        domain.User.ID.String(),
		Name:      domain.Name,
		NIPNumber: domain.NIP,
	}
}

func MapToBatchResponse(domains []admin.Domain) []Response {
	var responses []Response

	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return responses
}
