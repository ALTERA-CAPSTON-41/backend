package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	DOCTOR UserRole = "DOCTOR"
	NURSE  UserRole = "NURSE"
	ADMIN  UserRole = "ADMIN"
)

type User struct {
	gorm.Model
	ID       uuid.UUID
	Email    string `gorm:"unique"`
	Password string
	Role     UserRole `gorm:"type:enum('DOCTOR', 'ADMIN', 'NURSE')"`
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
		Role:     UserRole(request.Role),
	}
}

func MapToExistingUserModel(request UserRequest, id string) User {
	return User{
		ID:       uuid.MustParse(id),
		Email:    request.Email,
		Password: request.Password,
		Role:     UserRole(request.Role),
	}
}
