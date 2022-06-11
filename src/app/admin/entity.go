package admin

import (
	"time"

	"github.com/google/uuid"
)

type Domain struct {
	Name string
	NIP  string
	User UserReference
}

type UserReference struct {
	ID        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Services interface {
	GetAllAdmins() (admins []Domain, err error)
	GetAdminByID(id string) (admin *Domain, err error)
	CreateAdmin(admin Domain) (id string, err error)
	AmendAdminByID(id string, admin Domain) (err error)
	RemoveAdminByID(id string) (err error)
}

type Repositories interface {
	SelectAllData() (data []Domain, err error)
	SelectDataByID(id string) (data *Domain, err error)
	InsertData(data Domain) (id string, err error)
	UpdateByID(id string, data Domain) (err error)
	DeleteByID(id string) (err error)
}
