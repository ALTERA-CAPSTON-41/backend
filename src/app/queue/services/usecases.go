package services

import "clinic-api/src/app/queue"

type usecase struct {
	repo queue.Repositories
}

// CreateQueue implements queue.Services
func (uc *usecase) CreateQueue(queue queue.Domain) (string, error) {
	return uc.repo.InsertData(queue)
}

// GetAllQueues implements queue.Services
func (uc *usecase) GetAllQueues(polyclinic, from string) ([]queue.Domain, error) {
	return uc.repo.SelectAllData(polyclinic, from)
}

// AmendQueueByID implements queue.Services
func (uc *usecase) AmendQueueByID(id string, data queue.Domain) error {
	return uc.repo.UpdateByID(id, data)
}

// RemoveQueueByID implements queue.Services
func (uc *usecase) RemoveQueueByID(id string) error {
	return uc.repo.DeleteByID(id)
}

func NewService(repo queue.Repositories) queue.Services {
	return &usecase{repo}
}
