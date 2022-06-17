package nurse_repositories

import (
	"clinic-api/src/app/nurse"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Nurse struct {
	UserID       uuid.UUID `gorm:"primaryKey;size:191"`
	Name         string
	NIP          string `gorm:"column:nip"`
	SIP          string `gorm:"column:sip"`
	Address      string
	DOB          time.Time        `gorm:"type:date"`
	Gender       types.GenderEnum `gorm:"type:enum('MALE', 'FEMALE')"`
	PolyclinicID int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
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
	Email    string
	Password string
	Role     types.UserRoleEnum
}

func MapToDomain(record Nurse) nurse.Domain {
	return nurse.Domain{
		Name:    record.Name,
		NIP:     record.NIP,
		SIP:     record.SIP,
		Address: record.Address,
		DOB:     record.DOB,
		Gender:  record.Gender,
		Polyclinic: nurse.PolyclinicReference{
			ID:   record.Polyclinic.ID,
			Name: record.Polyclinic.Name,
		},
		User: nurse.UserReference{
			Email: record.User.Email,
		},
	}
}

func MapToNewRecord(domain nurse.Domain) Nurse {
	newUserID := uuid.Must(uuid.NewRandom())
	hashed, _ := utils.CreateHash(domain.User.Password)
	return Nurse{
		User: User{
			ID:       newUserID,
			Email:    domain.User.Email,
			Password: hashed,
			Role:     types.NURSE,
		},
		UserID:       newUserID,
		Name:         domain.Name,
		NIP:          domain.NIP,
		SIP:          domain.SIP,
		Address:      domain.Address,
		DOB:          domain.DOB,
		Gender:       domain.Gender,
		PolyclinicID: domain.Polyclinic.ID,
	}
}

func MapToExistingRecord(domain nurse.Domain) Nurse {
	return Nurse{
		Name:         domain.Name,
		NIP:          domain.NIP,
		SIP:          domain.SIP,
		Address:      domain.Address,
		DOB:          domain.DOB,
		Gender:       domain.Gender,
		PolyclinicID: domain.Polyclinic.ID,
	}
}

func MapToBatchDomain(records []Nurse) (domains []nurse.Domain) {
	for _, record := range records {
		domains = append(domains, MapToDomain(record))
	}
	return
}
