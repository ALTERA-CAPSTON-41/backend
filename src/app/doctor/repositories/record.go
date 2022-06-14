package repositories

import (
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Doctor struct {
	UserID       uuid.UUID `gorm:"primaryKey;size:191"`
	Name         string
	NIP          string `gorm:"column:nip"`
	SIP          string `gorm:"column:sip"`
	Address      string
	DOB          time.Time
	Gender       types.GenderEnum `gorm:"type:enum('MALE', 'FEMALE')"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	PolyclinicID int
	Polyclinic   Polyclinic
	User         User
}

type Polyclinic struct {
	ID   int
	Name string
}

type User struct {
	gorm.Model
	ID       uuid.UUID
	Email    string `gorm:"unique"`
	Password string
	Role     types.UserRoleEnum `gorm:"type:enum('DOCTOR', 'ADMIN', 'NURSE')"`
}
