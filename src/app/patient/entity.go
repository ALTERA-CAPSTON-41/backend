package patient

import (
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	ID        uuid.UUID
	Name      string
	NIK       string
	Phone     string
	Address   string
	DOB       time.Time
	Gender    types.GenderEnum
	BloodType string
}

type Services interface {
	HuntPatientByNameOrNIKOrAll(domain Domain, page int) ([]Domain, error)
	GetPatientByID(id string) (*Domain, error)
	CreatePatient(domain Domain) (id string, err error)
	AmendPatientByID(id string, domain Domain) error
	RemovePatientByID(id string) error
}

type Repositories interface {
	SearchDataByNameParam(name string, offset int) ([]Domain, error)
	SearchDataByNIKParam(nik string) ([]Domain, error)
	SelectAllData(offset int) ([]Domain, error)
	SelectDataByID(id string) (*Domain, error)
	UpdateByID(id string, domain Domain) error
	InsertData(domain Domain) (id string, err error)
	DeleteByID(id string) error
}
