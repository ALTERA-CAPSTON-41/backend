package request

import (
	"clinic-api/src/app/account"
	"clinic-api/src/types"
)

type UpdateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *UpdateRequest) MapToDomain() account.Domain {
	return account.Domain{
		Email:    req.Email,
		Password: req.Password,
		Role:     types.UserRoleEnum(req.Role),
	}
}

func (req *LoginRequest) MapToDomain() account.Domain {
	return account.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}
