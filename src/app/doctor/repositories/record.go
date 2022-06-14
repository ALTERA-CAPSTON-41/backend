package doctor_repositories

import (
	"clinic-api/src/app/doctor"
	"clinic-api/src/types"
	"clinic-api/src/utils"
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

func MapToDomain(record Doctor) doctor.Domain {
	return doctor.Domain{
		UserID:  record.User.ID,
		Name:    record.Name,
		NIP:     record.NIP,
		SIP:     record.NIP,
		Address: record.Address,
		DOB:     record.DOB,
		Gender:  record.Gender,
		Polyclinic: doctor.PolyclinicReference{
			ID:   record.Polyclinic.ID,
			Name: record.Polyclinic.Name,
		},
		User: doctor.UserReference{
			ID:        record.User.ID,
			Email:     record.User.Email,
			Password:  record.User.Password,
			Role:      record.User.Role,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		},
	}
}

func MapToNewRecord(domain doctor.Domain) Doctor {
	newUserID := uuid.Must(uuid.NewRandom())
	hashed, _ := utils.CreateHash(domain.User.Password)
	return Doctor{
		User: User{
			ID:       newUserID,
			Email:    domain.User.Email,
			Password: hashed,
			Role:     domain.User.Role,
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

func MapToExistingRecord(domain doctor.Domain) Doctor {
	return Doctor{
		UserID:       domain.User.ID,
		Name:         domain.Name,
		NIP:          domain.NIP,
		SIP:          domain.SIP,
		Address:      domain.Address,
		DOB:          domain.DOB,
		Gender:       domain.Gender,
		PolyclinicID: domain.Polyclinic.ID,
		User: User{
			ID:       domain.User.ID,
			Email:    domain.User.Email,
			Password: domain.User.Password,
			Role:     domain.User.Role,
		},
	}
}

func MapToBatchDomain(records []Doctor) []doctor.Domain {
	var domains []doctor.Domain

	for _, record := range records {
		domains = append(domains, MapToDomain(record))
	}
	return domains
}
