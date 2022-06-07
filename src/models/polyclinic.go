package models

import "gorm.io/gorm"

type Polyclinic struct {
	ID   int `gorm:"primaryKey"`
	Name string
	DeletedAt gorm.DeletedAt `gorm:"index"`
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

func MapToPolyclinicResponse(polyclinic Polyclinic) PolyclinicResponse {
	return PolyclinicResponse{
		ID: polyclinic.ID,
		Name: polyclinic.Name,
	}
}

func MapToPolyclinicBatchResponse(polyclinics []Polyclinic) []PolyclinicResponse {
	var response []PolyclinicResponse

	for _, polyclinic := range polyclinics {
		response = append(response, MapToPolyclinicResponse(polyclinic))
	}
	return response
}
