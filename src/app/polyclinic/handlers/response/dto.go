package response

import "clinic-api/src/app/polyclinic"

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ResponseWithStats struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	TotalDoctor int    `json:"total_doctor"`
	TotalNurse  int    `json:"total_nurse"`
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

func MapToResponseWithStats(domain polyclinic.Domain) ResponseWithStats {
	return ResponseWithStats{
		ID:          domain.ID,
		Name:        domain.Name,
		TotalDoctor: domain.TotalDoctor,
		TotalNurse:  domain.TotalNurse,
	}
}

func MapToBatchResponseWithStats(domains []polyclinic.Domain) (responses []ResponseWithStats) {
	for _, domain := range domains {
		responses = append(responses, MapToResponseWithStats(domain))
	}
	return
}
