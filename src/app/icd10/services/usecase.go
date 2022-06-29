package services

import "clinic-api/src/app/icd10"

type usecase struct {
	repo icd10.Repositories
}

// FindICD10ByCode implements icd10.Services
func (uc *usecase) FindICD10ByCode(code string) ([]icd10.Domain, error) {
	return uc.repo.SearchDataByCode(code)
}

func NewService(repo icd10.Repositories) icd10.Services {
	return &usecase{repo}
}
