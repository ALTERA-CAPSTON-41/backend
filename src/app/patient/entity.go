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

type Repositories interface {
	SearchDataByParams(name string, nik string) ([]Domain, error)
	SelectDataByID(id string) (*Domain, error)
	UpdateByID(id string, domain Domain) error
	InsertData(domain Domain) (id string, err error)
	DeleteByID(id string) error
}
