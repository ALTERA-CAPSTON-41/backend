package services

import "clinic-api/src/app/polyclinic"

type usecase struct {
	repo polyclinic.Repositories
}

// CreatePolyclinic implements polyclinic.Services
func (uc *usecase) CreatePolyclinic(polyclinic polyclinic.Domain) (int, error) {
	return uc.repo.InsertData(polyclinic)
}

// GetAllPolyclinics implements polyclinic.Services
func (uc *usecase) GetAllPolyclinics() ([]polyclinic.Domain, error) {
	return uc.repo.SelectAllData()
}

// GetPolyclinicByID implements polyclinic.Services
func (uc *usecase) GetPolyclinicByID(id int) (*polyclinic.Domain, error) {
	return uc.repo.SelectDataByID(id)
}

// AmendPolyclinicByID implements polyclinic.Services
func (uc *usecase) AmendPolyclinicByID(id int, polyclinic polyclinic.Domain) error {
	return uc.repo.UpdateByID(id, polyclinic)
}

// RemovePolyclinicByID implements polyclinic.Services
func (uc *usecase) RemovePolyclinicByID(id int) error {
	return uc.repo.DeleteByID(id)
}

func NewService(repo polyclinic.Repositories) polyclinic.Services {
	return &usecase{repo}
}
