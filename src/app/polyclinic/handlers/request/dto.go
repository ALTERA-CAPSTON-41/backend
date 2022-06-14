package request

import "clinic-api/src/app/polyclinic"

type Request struct {
	Name string `json:"name"`
}

func (request Request) MapToDomain() polyclinic.Domain {
	return polyclinic.Domain{Name: request.Name}
}
