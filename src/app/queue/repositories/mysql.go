package queue_repositories

import (
	"clinic-api/src/app/queue"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// InsertData implements queue.Repositories
func (repo *repository) InserData(data queue.Domain) (string, error) {
	var queueNumber int
	if err := repo.DB.Table("queues").Select("COALESCE(MAX(daily_queue_number), 0)").
		Where("daily_queue_date = CURDATE() AND polyclinic_id = ?", data.PolyclinicID).
		Find(&queueNumber).Error; err != nil {
		return uuid.Nil.String(), err
	}
	queueNumber++

	record := MapToNewRecord(data, queueNumber)
	if err := repo.DB.Create(&record).Error; err != nil {
		return uuid.Nil.String(), err
	}

	return record.ID.String(), nil
}

// SelectAllData implements queue.Repositories
func (repo *repository) SelectAllData(polyclinic, from string) ([]queue.Domain, error) {
	var (
		records      []Queue
		byPolyclinic string
		byDate       = " AND daily_queue_date = CURDATE()"
	)

	if polyclinic != "" {
		byPolyclinic = fmt.Sprint(" AND polyclinic_id = ", polyclinic)
	}

	if from != "" {
		byDate = fmt.Sprint(" AND daily_queue_date >= ", from)
	}

	if err := repo.DB.
		Preload("Patient").Preload("Polyclinic").
		Order("daily_queue_date DESC, daily_queue_number").
		Where("service_done_at IS NULL" + byPolyclinic + byDate).
		Find(&records).Error; err != nil {
		return nil, err
	}

	return MapToBatchDomain(records), nil
}

// UpdateByID implements queue.Repositories
func (repo *repository) UpdateByID(id string, data queue.Domain) error {
	record := MapToExistingRecord(id, data)
	alteration := repo.DB.Where("id", id).Updates(&record)
	if alteration.RowsAffected < 1 {
		return errors.New("record not found")
	}

	return alteration.Error
}

// DeleteByID implements queue.Repositories
func (repo *repository) DeleteByID(id string) error {
	deletion := repo.DB.Delete(new(Queue), id)
	if deletion.RowsAffected < 1 {
		return errors.New("record not found")
	}

	return deletion.Error
}

func NewMySQLRepository(DB *gorm.DB) queue.Repositories {
	return &repository{DB}
}
