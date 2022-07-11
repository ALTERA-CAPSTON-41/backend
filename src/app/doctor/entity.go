package doctor

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
	GetAllDoctors(polyclinic, page int) ([]Domain, error)
	GetDoctorByID(id string) (*Domain, error)
	CreateDoctor(doctor Domain) (string, error)
	AmendDoctorByID(id string, doctor Domain) error
	RemoveDoctorByID(id string) error
}

type Repositories interface {
	SelectAllData(polyclinic, offset int) ([]Domain, error)
	SelectDataByID(id string) (*Domain, error)
	LookupDataByEmail(email string) (string, error)
	InsertData(data Domain) (string, error)
	UpdateByID(id string, domain Domain) error
	DeleteByID(id string) error
	DeleteUserByID(id string) error
}
