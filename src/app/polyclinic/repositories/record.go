package repositories

import (
	"clinic-api/src/app/polyclinic"

	"gorm.io/gorm"
)

type Polyclinic struct {
	ID        int
	Name      string
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func MapToDomain(record Polyclinic) polyclinic.Domain {
	return polyclinic.Domain{
		Name: record.Name,
	}
}

func MapToNewRecord(domain polyclinic.Domain) Polyclinic {
	return Polyclinic{Name: domain.Name}
}

func MapToExistingRecord(id int, domain polyclinic.Domain) Polyclinic {
	return Polyclinic{
		ID:   id,
		Name: domain.Name,
	}
}

func MapToBatchDomain(records []Polyclinic) (domains []polyclinic.Domain) {
	for _, record := range records {
		domains = append(domains, MapToDomain(record))
	}
	return
}
