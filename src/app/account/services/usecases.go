package services

import (
	"clinic-api/src/app/account"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"errors"
)

type usecase struct {
	repo account.Repositories
}

// AttemptLogin implements account.Services
func (uc *usecase) AttemptLogin(domain account.Domain) (token string, role types.UserRoleEnum, err error) {
	var user *account.Domain

	if user, err = uc.repo.LookupAccountByEmail(domain.Email); err != nil {
		return "", "", err
	}

	if !utils.ValidateHash(domain.Password, user.Password) {
		return "", "", errors.New("password not match")
	}

	var account *account.UserDataDomain
	if user.Role == types.ADMIN {
		if account, err = uc.repo.LookupAdminByUserID(user.ID.String()); err != nil {
			return "", "", err
		}
	} else if user.Role == types.DOCTOR {
		if account, err = uc.repo.LookupDoctorByUserID(user.ID.String()); err != nil {
			return "", "", err
		}
	}

	token, err = utils.GenerateJwt(user.ID.String(), account.Name, account.NIP, user.Role)
	return token, user.Role, err
}

func NewService(repo account.Repositories) account.Services {
	return &usecase{repo}
}
