package queue_repositories

import (
	"clinic-api/src/app/queue"
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Queue struct {
	gorm.Model
	ID               uuid.UUID `gorm:"size:191"`
	PatientID        uuid.UUID
	PolyclinicID     int
	PatientStatus    types.PatientStatusEnum `gorm:"type:enum('OUTPATIENT', 'REFERRED')"`
	DailyQueueNumber int
	DailyQueueDate   time.Time `gorm:"type:date"`
	ServiceDoneAt    time.Time `gorm:"default:null"`
	Patient          Patient
	Polyclinic       Polyclinic
}

type Patient struct {
	gorm.Model
	ID        uuid.UUID
	Name      string
	NIK       string
	Phone     string
	Address   string
	DOB       time.Time
	Gender    types.GenderTypeEnum
	BloodType string
}

type Polyclinic struct {
	ID   int
	Name string
}

func MapToDomain(record Queue) queue.Domain {
	return queue.Domain{
		ID: record.ID,
		Patient: queue.PatientReference{
			ID:     record.Patient.ID,
			Name:   record.Patient.Name,
			Gender: record.Patient.Gender,
		},
		Polyclinic: queue.PolyclinicReference{
			ID: record.Polyclinic.ID, Name: record.Polyclinic.Name,
		},
		PatientStatus:    record.PatientStatus,
		DailyQueueNumber: record.DailyQueueNumber,
		DailyQueueDate:   record.DailyQueueDate,
		ServiceDoneAt:    record.ServiceDoneAt,
	}
}

func MapToNewRecord(domain queue.Domain, queueNumber int) Queue {
	return Queue{
		ID:               uuid.Must(uuid.NewRandom()),
		PatientID:        domain.PatientID,
		PolyclinicID:     domain.PolyclinicID,
		PatientStatus:    domain.PatientStatus,
		DailyQueueNumber: queueNumber,
		DailyQueueDate:   domain.DailyQueueDate,
	}
}

func MapToExistingRecord(id string, domain queue.Domain) Queue {
	return Queue{
		ID:             uuid.MustParse(id),
		PatientID:      domain.PatientID,
		PolyclinicID:   domain.PolyclinicID,
		PatientStatus:  domain.PatientStatus,
		DailyQueueDate: domain.DailyQueueDate,
		ServiceDoneAt:  domain.ServiceDoneAt,
	}
}

func MapToBatchDomain(records []Queue) (domains []queue.Domain) {
	for _, record := range records {
		domains = append(domains, MapToDomain(record))
	}
	return
}
