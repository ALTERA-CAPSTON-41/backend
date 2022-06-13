package models

import (
	"clinic-api/src/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Queue struct {
	gorm.Model
	ID               uuid.UUID `gorm:"size:191"`
	PatientID        uuid.UUID
	PolyclinicID     int
	PatientStatus    PatientStatus `gorm:"type:enum('OUTPATIENT', 'REFERRED')"`
	DailyQueueNumber int
	DailyQueueDate   time.Time `gorm:"type:date"`
	ServiceDoneAt    time.Time `gorm:"default:null"`
	Patient          Patient
	Polyclinic       Polyclinic
}

type QueueRequest struct {
	PatientID      string `json:"patient_id"`
	PolyclinicID   int    `json:"polyclinic_id"`
	PatientStatus  string `json:"patient_status"`
	DailyQueueDate string `json:"daily_queue_date"`
	ServiceDoneAt  string `json:"service_done_at"`
}

type QueueResponse struct {
	ID               uuid.UUID          `json:"id"`
	Patient          PatientResponse    `json:"patient"`
	Polyclinic       PolyclinicResponse `json:"polyclinic"`
	PatientStatus    PatientStatus      `json:"patient_status"`
	DailyQueueNumber int                `json:"daily_queue_number"`
	DailyQueueDate   string             `json:"daily_queue_date"`
	ServiceDoneAt    string             `json:"service_done_at"`
}

func MapToNewQueue(request QueueRequest, queueNumber int) Queue {
	return Queue{
		ID:               uuid.Must(uuid.NewRandom()),
		PatientID:        uuid.MustParse(request.PatientID),
		PolyclinicID:     request.PolyclinicID,
		PatientStatus:    PatientStatus(request.PatientStatus),
		DailyQueueNumber: queueNumber,
		DailyQueueDate:   utils.ConvertStringToDate(request.DailyQueueDate),
	}
}

func MapToExistingQueue(request QueueRequest, id string) Queue {
	return Queue{
		ID:             uuid.MustParse(id),
		PatientID:      uuid.MustParse(request.PatientID),
		PolyclinicID:   request.PolyclinicID,
		PatientStatus:  PatientStatus(request.PatientStatus),
		DailyQueueDate: utils.ConvertStringToDate(request.DailyQueueDate),
		ServiceDoneAt:  utils.ConvertStringToDatetime(request.ServiceDoneAt),
	}
}

func MapToQueueResponse(queue Queue) QueueResponse {
	serviceDoneAt := utils.ConvertDatetimeToString(queue.ServiceDoneAt)
	if serviceDoneAt == "0001-01-01 00:00:00" {
		serviceDoneAt = ""
	}

	return QueueResponse{
		ID: queue.ID,
		Patient: PatientResponse{
			ID:        queue.Patient.ID,
			Name:      queue.Patient.Name,
			NIK:       queue.Patient.NIK,
			Phone:     queue.Patient.Phone,
			Address:   queue.Patient.Address,
			DOB:       utils.ConvertDateToString(queue.Patient.DOB),
			Gender:    queue.Patient.Gender,
			BloodType: queue.Patient.BloodType,
		},
		Polyclinic: PolyclinicResponse{
			ID: queue.Polyclinic.ID, Name: queue.Polyclinic.Name,
		},
		PatientStatus:    queue.PatientStatus,
		DailyQueueNumber: queue.DailyQueueNumber,
		DailyQueueDate:   utils.ConvertDateToString(queue.DailyQueueDate),
		ServiceDoneAt:    serviceDoneAt,
	}
}

func MapToQueueBatchResponse(queues []Queue) []QueueResponse {
	var response []QueueResponse

	for _, queue := range queues {
		response = append(response, MapToQueueResponse(queue))
	}
	return response
}
