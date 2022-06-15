package models

import (
	"clinic-api/src/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID
	Email    string `gorm:"unique"`
	Password string
	Role     types.UserRoleEnum `gorm:"type:enum('DOCTOR', 'ADMIN', 'NURSE')"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type AuthErrorResponse struct {
	Reason string `json:"reason"`
}

func MapToNewUserModel(request UserRequest) User {
	return User{
		ID:       uuid.Must(uuid.NewRandom()),
		Email:    request.Email,
		Password: request.Password,
		Role:     types.UserRoleEnum(request.Role),
	}
}

func MapToExistingUserModel(request UserRequest, id string) User {
	return User{
		ID:       uuid.MustParse(id),
		Email:    request.Email,
		Password: request.Password,
		Role:     types.UserRoleEnum(request.Role),
	}
}
