package services

import (
	"clinic-api/src/app/account"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"errors"
	"fmt"
)

type usecase struct {
	repo account.Repositories
}

// AttemptLogin implements account.Services
func (uc *usecase) AttemptLogin(domain account.Domain) (*account.UserDataDomain, error) {
	var user *account.Domain
	var err error

	if user, err = uc.repo.LookupAccountByEmail(domain.Email); err != nil {
		return nil, err
	}

	if !utils.ValidateHash(domain.Password, user.Password) {
		fmt.Println("result", domain.Password, user.Password, utils.ValidateHash(domain.Password, user.Password))
		return nil, errors.New("password not match")
	}

	var result *account.UserDataDomain
	if user.Role == types.ADMIN {
		if result, err = uc.repo.LookupAdminByUserID(user.ID.String()); err != nil {
			return nil, err
		}
	} else if user.Role == types.DOCTOR {
		if result, err = uc.repo.LookupDoctorByUserID(user.ID.String()); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func NewService(repo account.Repositories) account.Services {
	return &usecase{repo}
}
