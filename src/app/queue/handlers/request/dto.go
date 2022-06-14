package request

import (
	"clinic-api/src/app/queue"
	"clinic-api/src/types"
	"clinic-api/src/utils"

	"github.com/google/uuid"
)

type NewRequest struct {
	PatientID      string `json:"patient_id"`
	PolyclinicID   int    `json:"polyclinic_id"`
	PatientStatus  string `json:"patient_status"`
	DailyQueueDate string `json:"daily_queue_date"`
}

type UpdateRequest struct {
	PatientID      string `json:"patient_id"`
	PolyclinicID   int    `json:"polyclinic_id"`
	PatientStatus  string `json:"patient_status"`
	DailyQueueDate string `json:"daily_queue_date"`
	ServiceDoneAt  string `json:"service_done_at"`
}

func (req *NewRequest) MapToDomain() queue.Domain {
	return queue.Domain{
		PatientID:      uuid.MustParse(req.PatientID),
		PolyclinicID:   req.PolyclinicID,
		PatientStatus:  types.PatientStatusEnum(req.PatientStatus),
		DailyQueueDate: utils.ConvertStringToDate(req.DailyQueueDate),
	}
}

func (req *UpdateRequest) MapToDomain() queue.Domain {
	return queue.Domain{
		PatientID:      uuid.MustParse(req.PatientID),
		PolyclinicID:   req.PolyclinicID,
		PatientStatus:  types.PatientStatusEnum(req.PatientStatus),
		DailyQueueDate: utils.ConvertStringToDate(req.DailyQueueDate),
		ServiceDoneAt:  utils.ConvertStringToDatetime(req.ServiceDoneAt),
	}
}
