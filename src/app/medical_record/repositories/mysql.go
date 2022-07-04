package medicalrecord_repositories

import (
	medicalrecord "clinic-api/src/app/medical_record"
	"clinic-api/src/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// InsertData implements medicalrecord.Repositories
func (repo *repository) InsertData(domain medicalrecord.Domain) (id string, err error) {
	record := MapToNewRecord(domain)

	if err := repo.DB.Create(&record).Error; err != nil {
		return "", err
	}
	return record.ID.String(), nil
}

// LookupICD10Data implements medicalrecord.Repositories
func (repo *repository) LookupICD10Data(icd10Code string) (ICD10Description string, err error) {
	endpoint := fmt.Sprintf("http://icd10api.com/?code=%s&desc=long&r=json", icd10Code)
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	var body ICDResponse
	json.NewDecoder(resp.Body).Decode(&body)

	return body.Description, nil
}

// SelectDataByID implements medicalrecord.Repositories
func (repo *repository) SelectDataByID(id string) (*medicalrecord.Domain, error) {
	var record MedicalRecord

	if err := repo.DB.Where("ID = ?", id).
		Preload("Polyclinic").
		Preload("Patient").
		Preload("Doctor").
		First(&record).Error; err != nil {
		return nil, err
	}
	result := record.MapToDomain()
	result.Patient.Age = utils.CountIntervalByYearRoundDown(result.Patient.DOB, result.CreatedAt)
	return &result, nil
}

// SelectDataByPatientNIK implements medicalrecord.Repositories
func (repo *repository) SelectDataByPatientNIK(nik string) ([]medicalrecord.Domain, error) {
	var records []MedicalRecord

	if err := repo.DB.
		Joins("JOIN patients ON patients.nik = ?", nik).Preload("Patient").
		Preload("Doctor").Preload("Polyclinic").
		Find(&records).Error; err != nil {
		return nil, err
	}
	return MapToBatchDomain(records), nil
}

func NewMySQLRepository(conn *gorm.DB) medicalrecord.Repositories {
	return &repository{conn}
}
