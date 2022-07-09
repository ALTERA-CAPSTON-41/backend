package medicalrecord_repositories

import (
	medicalrecord "clinic-api/src/app/medical_record"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// InsertData implements medicalrecord.Repositories
func (repo *repository) InsertData(domain medicalrecord.Domain) (id string, err error) {
	// set default value of patient status
	if strings.ToUpper(string(domain.PatientStatus)) != string(types.OUTPATIENT) ||
		strings.ToUpper(string(domain.PatientStatus)) != string(types.REFERRED) {
		domain.PatientStatus = types.OUTPATIENT
	}
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

// SelectPatientIDByNIK implements medicalrecord.Repositories
func (repo *repository) SelectPatientIDByNIK(nik string) (string, error) {
	var id string
	if err := repo.DB.Model(&Patient{}).Select("id").First(&id, "nik = ?", nik).Error; err != nil {
		return "", err
	}

	return id, nil
}

// SelectDataByPatientID implements medicalrecord.Repositories
func (repo *repository) SelectDataByPatientID(id string) ([]medicalrecord.Domain, error) {
	var records []MedicalRecord

	if err := repo.DB.
		Preload("Patient").
		Preload("Doctor").Preload("Polyclinic").
		Find(&records, "patient_id = ?", id).Error; err != nil {
		return nil, err
	}
	return MapToBatchDomain(records), nil
}

func NewMySQLRepository(conn *gorm.DB) medicalrecord.Repositories {
	return &repository{conn}
}
