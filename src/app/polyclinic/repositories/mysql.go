package polyclinic_repositories

import (
	"clinic-api/src/app/polyclinic"
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// InserData implements polyclinic.Repositories
func (repo *repository) InsertData(data polyclinic.Domain) (int, error) {
	record := MapToNewRecord(data)
	if err := repo.DB.Create(&record).Error; err != nil {
		return 0, err
	}

	return record.ID, nil
}

// SelectAllData implements polyclinic.Repositories
func (repo *repository) SelectAllData() ([]polyclinic.Domain, error) {
	var records []Polyclinic
	if err := repo.DB.Find(&records).Error; err != nil {
		return nil, err
	}

	return MapToBatchDomain(records), nil
}

// SelectDataByID implements polyclinic.Repositories
func (repo *repository) SelectDataByID(id int) (*polyclinic.Domain, error) {
	var record Polyclinic
	if err := repo.DB.First(&record, id).Error; err != nil {
		return nil, err
	}

	polyclinic := MapToDomain(record)
	return &polyclinic, nil
}

// UpdateByID implements polyclinic.Repositories
func (repo *repository) UpdateByID(id int, data polyclinic.Domain) error {
	record := MapToExistingRecord(id, data)
	alteration := repo.DB.Where("id", id).Updates(&record)
	if alteration.RowsAffected < 1 && alteration.Error == nil {
		return errors.New("record not found")
	}

	return alteration.Error
}

// DeleteByID implements polyclinic.Repositories
func (repo *repository) DeleteByID(id int) error {
	deletion := repo.DB.Delete(new(Polyclinic), id)
	if deletion.RowsAffected < 1 && deletion.Error == nil {
		return errors.New("record not found")
	}

	return deletion.Error
}

// CountDoctorByPolyclinic implements polyclinic.Repositories
func (repo *repository) CountDoctorByPolyclinic(polyclinic int) (int, error) {
	var total int64
	err := repo.DB.Table("doctors").Where("polyclinic_id = ? AND deleted_at IS NULL", polyclinic).Count(&total).Error
	return int(total), err
}

// CountNurseByPolyclinic implements polyclinic.Repositories
func (repo *repository) CountNurseByPolyclinic(polyclinic int) (int, error) {
	var total int64
	err := repo.DB.Table("nurses").Where("polyclinic_id = ? AND deleted_at IS NULL", polyclinic).Count(&total).Error
	return int(total), err
}

func NewMySQLRepository(DB *gorm.DB) polyclinic.Repositories {
	return &repository{DB}
}
