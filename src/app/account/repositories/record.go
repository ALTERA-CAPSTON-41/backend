package account_repositories

import (
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
