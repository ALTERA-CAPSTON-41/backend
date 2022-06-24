package dashboard_repositories

import (
	"clinic-api/src/app/dashboard"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// CountData implements dashboard.Repositories
func (repo *repository) CountData(table string) (int, error) {
	var total int64
	err := repo.DB.Table(table).Where("deleted_at IS NULL").Count(&total).Error
	return int(total), err
}

// CountQueueData implements dashboard.Repositories
func (repo *repository) CountQueueData() (int, error) {
	var total int64
	err := repo.DB.
		Table("queues").
		Where("deleted_at IS NULL AND daily_queue_date = CURDATE() AND service_done_at IS NULL").
		Count(&total).Error
	return int(total), err
}

func NewMySQLRepository(DB *gorm.DB) dashboard.Repositories {
	return &repository{DB}
}
