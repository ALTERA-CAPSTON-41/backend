package nurse

import (
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
)

type Domain struct {
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

type Services interface {
	CreateNurse(nurse Domain) (*Domain, error)
	GetAllNurses() ([]Domain, error)
	GetNurseByID(id string) (*Domain, error)
	AmendNurseByID(id string, doctor Domain) error
	RemoveNurseByID(id string)
}

type Repositories interface {
	InsertData(data Domain) (string, error)
	SelectAllData() ([]Domain, error)
	SelectDataByID(id string) (*Domain, error)
	UpdateByID(id string, domain Domain) error
	DeleteByID(id string) error
}
