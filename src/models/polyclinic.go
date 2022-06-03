package models

type Polyclinic struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type PolyclinicRequest struct {
	Name string `json:"name"`
}

type PolyclinicResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func MapToNewPolyclinic(request PolyclinicRequest) Polyclinic {
	return Polyclinic{
		Name: request.Name,
	}
}

func MapToExistingPolyclinic(request PolyclinicRequest, id int) Polyclinic {
	return Polyclinic{
		ID:   id,
		Name: request.Name,
	}
}

func MapToPolyclinic(polyclinic Polyclinic) PolyclinicResponse {
	return PolyclinicResponse(polyclinic)
}

func MapToPolyclinicBatch(polyclinics []Polyclinic) []PolyclinicResponse {
	var response []PolyclinicResponse

	for _, polyclinic := range polyclinics {
		response = append(response, PolyclinicResponse(polyclinic))
	}
	return response
}
