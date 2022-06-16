package services

import (
	"clinic-api/src/app/account"
)

type usecase struct {
	repo account.Repositories
}

// AttemptLogin implements account.Services
func (uc *usecase) AttemptLogin(domain account.Domain) (*account.UserDataDomain, error) {
	panic("unimplemented")
}

func NewService(repo account.Repositories) account.Services {
	return &usecase{repo}
}
