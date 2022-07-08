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

// GetAllPolyclinicsWithStats implements polyclinic.Services
func (uc *usecase) GetAllPolyclinicsWithStats() ([]polyclinic.Domain, error) {
	polyclinics, err := uc.repo.SelectAllData()
	if err != nil {
		return nil, err
	}
	for index, polyclinic := range polyclinics {
		var totalDoctor, totalNurse int
		totalDoctor, err := uc.repo.CountDoctorByPolyclinic(polyclinic.ID)
		if err != nil {
			return nil, err
		}

		totalNurse, err = uc.repo.CountNurseByPolyclinic(polyclinic.ID)
		if err != nil {
			return nil, err
		}

		polyclinics[index].TotalDoctor = totalDoctor
		polyclinics[index].TotalNurse = totalNurse
	}

	return polyclinics, nil
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
