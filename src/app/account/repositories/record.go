package account_repositories

import (
	"clinic-api/src/app/account"
	"clinic-api/src/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID
	Email    string `gorm:"unique"`
	Password string
	Role     types.UserRoleEnum `gorm:"type:enum('DOCTOR', 'ADMIN', 'NURSE')"`
}

type Doctor struct {
	UserID uuid.UUID `gorm:"primaryKey;size:191"`
	Name   string
	NIP    string `gorm:"column:nip"`
}

type Admin struct {
	UserID uuid.UUID `gorm:"primaryKey;size:191"`
	Name   string
	NIP    string `gorm:"column:nip"`
}

type Nurse struct {
	UserID uuid.UUID `gorm:"primaryKey;size:191"`
	Name   string
	NIP    string `gorm:"column:nip"`
}

func (record *User) MapToDomain() account.Domain {
	return account.Domain{
		ID:       record.ID,
		Email:    record.Email,
		Password: record.Password,
		Role:     record.Role,
	}
}

func (record *Admin) MapToDomain() account.UserDataDomain {
	return account.UserDataDomain{
		ID:   record.UserID,
		Name: record.Name,
		NIP:  record.NIP,
		Role: types.ADMIN,
	}
}

func (record *Doctor) MapToDomain() account.UserDataDomain {
	return account.UserDataDomain{
		ID:   record.UserID,
		Name: record.Name,
		NIP:  record.NIP,
		Role: types.DOCTOR,
	}
}

func (record *Nurse) MapToDomain() account.UserDataDomain {
	return account.UserDataDomain{
		ID:   record.UserID,
		Name: record.Name,
		NIP:  record.NIP,
		Role: types.NURSE,
	}
}
