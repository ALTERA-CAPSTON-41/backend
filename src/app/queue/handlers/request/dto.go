package request

import (
	"clinic-api/src/app/queue"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"time"

	"github.com/google/uuid"
)

type NewRequest struct {
	PatientID     string `json:"patient_id"`
	PolyclinicID  int    `json:"polyclinic_id"`
	PatientStatus string `json:"patient_status"`
}

type UpdateRequest struct {
	ServiceDoneAt string `json:"service_done_at"`
}

func (req *NewRequest) MapToDomain() queue.Domain {
	return queue.Domain{
		PatientID:      uuid.MustParse(req.PatientID),
		PolyclinicID:   req.PolyclinicID,
		PatientStatus:  types.PatientStatusEnum(req.PatientStatus),
		DailyQueueDate: utils.ConvertStringToDate(time.Now().Format("2006-01-02")),
	}
}

func (req *UpdateRequest) MapToDomain() queue.Domain {
	return queue.Domain{
		ServiceDoneAt: utils.ConvertStringToDatetime(req.ServiceDoneAt),
	}
}
