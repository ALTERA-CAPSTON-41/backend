package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	ID        uuid.UUID      `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	NIK       string         `json:"nik"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	DOB       time.Time      `json:"dob"`
	Gender    string         `json:"gender"`
	BloodType string         `json:"blood_type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
