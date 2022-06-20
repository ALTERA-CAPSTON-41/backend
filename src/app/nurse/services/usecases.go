package services

import "clinic-api/src/app/nurse"

type usecase struct {
	repo nurse.Repositories
}

// CreateNurse implements nurse.Services
func (uc *usecase) CreateNurse(nurse nurse.Domain) (string, error) {
	return uc.repo.InsertData(nurse)
}

// GetAllNurses implements nurse.Services
func (uc *usecase) GetAllNurses() ([]nurse.Domain, error) {
	return uc.repo.SelectAllData()
}

// GetNurseByID implements nurse.Services
func (uc *usecase) GetNurseByID(id string) (*nurse.Domain, error) {
	return uc.repo.SelectDataByID(id)
}

// AmendNurseByID implements nurse.Services
func (uc *usecase) AmendNurseByID(id string, nurse nurse.Domain) error {
	return uc.repo.UpdateByID(id, nurse)
}

// RemoveNurseByID implements nurse.Services
func (uc *usecase) RemoveNurseByID(id string) error {
	if err := uc.repo.DeleteByID(id); err != nil {
		return err
	}

	return uc.repo.DeleteUserByID(id)
}

func NewService(repo nurse.Repositories) nurse.Services {
	return &usecase{repo}
}
