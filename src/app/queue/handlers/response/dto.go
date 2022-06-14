package response

import (
	"clinic-api/src/app/queue"
	"clinic-api/src/utils"
)

type Response struct {
	ID               string             `json:"id"`
	Patient          PatientResponse    `json:"patient"`
	Polyclinic       PolyclinicResponse `json:"polyclinic"`
	PatientStatus    string             `json:"patient_status"`
	DailyQueueNumber int                `json:"daily_queue_number"`
	DailyQueueDate   string             `json:"daily_queue_date"`
	ServiceDoneAt    string             `json:"service_done_at"`
}

type PatientResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type PolyclinicResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateResponse struct {
	ID string `json:"id"`
}

func MapToResponse(domain queue.Domain) Response {
	return Response{
		ID: domain.ID.String(),
		Patient: PatientResponse{
			ID:     domain.Patient.ID.String(),
			Name:   domain.Patient.Name,
			Gender: string(domain.Patient.Gender),
		},
		Polyclinic: PolyclinicResponse{
			ID: domain.Polyclinic.ID, Name: domain.Polyclinic.Name,
		},
		PatientStatus:    string(domain.PatientStatus),
		DailyQueueNumber: domain.DailyQueueNumber,
		DailyQueueDate:   utils.ConvertDateToString(domain.DailyQueueDate),
		ServiceDoneAt:    utils.ConvertDatetimeToString(domain.ServiceDoneAt),
	}
}

func MapToBatchResponse(domains []queue.Domain) (responses []Response) {
	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return
}
