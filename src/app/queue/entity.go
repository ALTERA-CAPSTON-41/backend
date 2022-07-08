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
	Gender types.GenderEnum
}

type PolyclinicReference struct {
	ID   int
	Name string
}

type Services interface {
	CreateQueue(queue Domain) (string, error)
	GetAllQueues(fromDate string, polyclinic, page int) ([]Domain, error)
	AmendQueueByID(id string, queue Domain) error
	RemoveQueueByID(id string) error
}

type Repositories interface {
	SelectQueueNumber(polyclinicID int) (int, error)
	InsertData(data Domain) (string, error)
	SelectAllData(fromDate string, polyclinic, offset int) ([]Domain, error)
	UpdateByID(id string, data Domain) error
	DeleteByID(id string) error
}
