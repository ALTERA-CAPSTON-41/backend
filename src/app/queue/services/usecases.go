package services

import "clinic-api/src/app/queue"

type usecase struct {
	repo queue.Repositories
}

// CreateQueue implements queue.Services
func (uc *usecase) CreateQueue(queue queue.Domain) (string, error) {
	panic("unimplemented")
}

// GetAllQueues implements queue.Services
func (uc *usecase) GetAllQueues(polyclinic, from string) ([]queue.Domain, error) {
	panic("unimplemented")
}

// AmendQueueByID implements queue.Services
func (uc *usecase) AmendQueueByID(id string, data queue.Domain) error {
	panic("unimplemented")
}

// RemoveQueueByID implements queue.Services
func (uc *usecase) RemoveQueueByID(id string) error {
	panic("unimplemented")
}

func NewService(repo queue.Repositories) queue.Services {
	return &usecase{repo}
}
