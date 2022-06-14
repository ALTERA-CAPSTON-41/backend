package response

import "clinic-api/src/app/polyclinic"

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateResponse struct {
	ID int `json:"id"`
}

func MapToResponse(domain polyclinic.Domain) Response {
	return Response{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func MapToBatchResponse(domains []polyclinic.Domain) (responses []Response) {
	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return
}
