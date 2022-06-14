package doctor

import (
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	UserID     uuid.UUID
	Name       string
	NIP        string
	SIP        string
	Address    string
	DOB        time.Time
	Gender     types.GenderEnum
	Polyclinic PolyclinicReference
	User       UserReference
}

type PolyclinicReference struct {
	ID   int
	Name string
}

type UserReference struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Role      types.UserRoleEnum
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repositories interface {
	SelectAllData() ([]Domain, error)
	SelectDataByID(id string) (*Domain, error)
	InsertData(data Domain) (string, error)
	UpdateByID(id string, domain Domain) error
	DeleteByID(id string) error
}
