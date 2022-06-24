package services

import (
	"clinic-api/src/app/dashboard"
	"errors"
	"fmt"
)

type usecase struct {
	repo dashboard.Repositories
}

func (uc *usecase) GetTotal(feature string) (int, error) {
	if feature == "patient" || feature == "polyclinic" ||
		feature == "doctor" || feature == "nurse" {
		table := fmt.Sprintf("%ss", feature)
		return uc.repo.CountData(table)
	}

	if feature == "queue" {
		return uc.repo.CountQueueData()
	}

	return 0, errors.New("forbidden")
}

func NewService(repo dashboard.Repositories) dashboard.Services {
	return &usecase{repo}
}
