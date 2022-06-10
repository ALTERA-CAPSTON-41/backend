package admin_repositories

import (
	"clinic-api/src/app/admin"
	"clinic-api/src/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	UserID    uuid.UUID `gorm:"primaryKey;size:191"`
	Name      string
	NIP       string `gorm:"column:nip"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	User      User
}

type User struct {
	gorm.Model
	ID       uuid.UUID
	Email    string
	Password string
}

func MapToDomain(record Admin) admin.Domain {
	return admin.Domain{
		Name: record.Name,
		NIP:  record.NIP,
		User: admin.UserReference{
			ID:        record.User.ID,
			Email:     record.User.Email,
			Password:  record.User.Password,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		},
	}
}

func MapToNewRecord(domain admin.Domain) Admin {
	newUserID := uuid.Must(uuid.NewRandom())
	hashed, _ := utils.CreateHash(domain.User.Password)
	return Admin{
		UserID: newUserID,
		Name:   domain.Name,
		NIP:    domain.NIP,
		User: User{
			ID:       newUserID,
			Email:    domain.User.Email,
			Password: hashed,
		},
	}
}

func MapToExistingRecord(domain admin.Domain) Admin {
	return Admin{
		UserID: domain.User.ID,
		Name:   domain.Name,
		NIP:    domain.NIP,
		User: User{
			ID:       domain.User.ID,
			Email:    domain.User.Email,
			Password: domain.User.Password,
		},
	}
}

func MapToBatchDomain(records []Admin) []admin.Domain {
	var domains []admin.Domain

	for _, record := range records {
		domains = append(domains, MapToDomain(record))
	}
	return domains
}
