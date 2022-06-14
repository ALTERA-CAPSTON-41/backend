package queue

import (
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID               uuid.UUID
	PatientID        uuid.UUID
	PolyclinicID     int
	PatientStatus    types.PatientStatusEnum
	DailyQueueNumber int
	DailyQueueDate   time.Time
	ServiceDoneAt    time.Time
	Patient          PatientReference
	Polyclinic       PolyclinicReference
}

type PatientReference struct {
	ID     uuid.UUID
	Name   string
	Gender types.GenderTypeEnum
}

type PolyclinicReference struct {
	ID   int
	Name string
}

type Services interface {
	CreateQueue(queue Domain) (string, error)
	GetAllQueues() ([]Domain, error)
	AmendPolyclinicByID(id string, queue Domain) error
	RemovePolyclinicByID(id string)
}

type Repositories interface {
	InserData(data Domain) (string, error)
	SelectAllData() ([]Domain, error)
	UpdateByID(id string, data Domain) error
	DeleteByID(id string) error
}
