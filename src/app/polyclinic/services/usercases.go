package services

import "clinic-api/src/app/polyclinic"

type usecase struct {
	repo polyclinic.Repositories
}

func (uc *usecase) CreatePolyclinic(polyclinic polyclinic.Domain) (int, error) {
	panic("unimplemented")
}

func (uc *usecase) GetAllPolyclinics() ([]polyclinic.Domain, error) {
	panic("unimplemented")
}

func (uc *usecase) GetPolyclinicByID(id int) (*polyclinic.Domain, error) {
	panic("unimplemented")
}

func (uc *usecase) AmendPolyclinicByID(id int, polyclinic polyclinic.Domain) error {
	panic("unimplemented")
}

func (uc *usecase) RemovePolyclinicByID(id int) error {
	panic("unimplemented")
}

func NewService(repo polyclinic.Repositories) polyclinic.Services {
	return &usecase{repo}
}
