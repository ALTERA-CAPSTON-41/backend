package response

import "clinic-api/src/app/icd10"

type Response struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	Reason string `json:"reason"`
}

func MapToResponse(domain icd10.Domain) Response {
	return Response{
		Name:        domain.Name,
		Description: domain.Description,
	}
}

func MapToBatchResponse(domains []icd10.Domain) (responses []Response) {
	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return
}
