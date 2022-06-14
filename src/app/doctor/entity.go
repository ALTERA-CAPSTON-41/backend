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
	CreatedAt time.Time
	UpdatedAt time.Time
}
