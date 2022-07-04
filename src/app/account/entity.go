package account

import (
	"clinic-api/src/types"

	"github.com/google/uuid"
)

type Domain struct {
	ID       uuid.UUID
	Email    string
	Password string
	Role     types.UserRoleEnum
}

type UserDataDomain struct {
	ID   uuid.UUID
	Name string
	NIP  string
	Role types.UserRoleEnum
}

type Services interface {
	AttemptLogin(domain Domain) (token string, role types.UserRoleEnum, err error)
}

type Repositories interface {
	LookupAccountByEmail(email string) (*Domain, error)
	LookupDoctorByUserID(id string) (*UserDataDomain, error)
	LookupAdminByUserID(id string) (*UserDataDomain, error)
	LookupNurseByUserID(id string) (*UserDataDomain, error)
}
