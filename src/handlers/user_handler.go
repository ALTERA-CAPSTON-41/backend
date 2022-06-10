package handlers

import (
	"clinic-api/src/models"
	"clinic-api/src/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AttemptLoginUser(c echo.Context) error {
	var userRequest models.LoginRequest
	if err := c.Bind(&userRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	userData := lookupUserByEmail(userRequest.Email)
	if userData.ID == uuid.Nil {
		return utils.CreateEchoResponse(c, http.StatusUnauthorized, models.AuthErrorResponse{
			Reason: "email not registered yet",
		})
	}

	if utils.ValidateHash(userRequest.Password, userData.Password) {
		doctorUser := lookupDoctorFromUserData(userData.ID.String())
		token, _ := utils.GenerateJwt(
			userData.ID.String(),
			doctorUser.Name,
			doctorUser.NIP,
			utils.UserRole(userData.Role),
		)
		utils.SetJwtCookie(c, token)
		return utils.CreateEchoResponse(c, http.StatusCreated, models.AuthResponse{
			Token: token,
		})
	}

	return utils.CreateEchoResponse(c, http.StatusUnauthorized, models.AuthErrorResponse{
		Reason: "password incorrect",
	})
}

// in controller helper
func lookupUserByEmail(email string) models.User {
	var userData models.User
	models.DB.Where("email = ?", email).First(&userData)
	return userData
}

func lookupDoctorFromUserData(userID string) models.Doctor {
	var doctor models.Doctor
	models.DB.Where("user_id = ?", userID).First(&doctor)
	return doctor
}
