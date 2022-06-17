package patient_repositories

import (
	"clinic-api/src/types"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey;size:191"`
	Name      string
	NIK       string
	Phone     string
	Address   string
	DOB       time.Time
	Gender    types.GenderEnum `gorm:"type:enum('MALE', 'FEMALE')"`
	BloodType string
}
