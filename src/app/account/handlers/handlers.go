package account_handlers

import (
	"clinic-api/src/app/account"
	"clinic-api/src/app/account/handlers/request"
	"clinic-api/src/app/account/handlers/response"
	"clinic-api/src/types"
	"clinic-api/src/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services account.Services
}

// onAttemptLogin
func (h *Handler) AttemptLoginHandler(c echo.Context) error {
	var loginRequest request.LoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	var token string
	var role types.UserRoleEnum
	var err error
	if token, role, err = h.services.AttemptLogin(loginRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusUnauthorized, response.AuthErrorResponse{
				Reason: "email not registered yet",
			})
		}

		if strings.Contains(err.Error(), "password not match") {
			return utils.CreateEchoResponse(c, http.StatusUnauthorized, response.AuthErrorResponse{
				Reason: "password entered incorrect",
			})
		}

		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	utils.SetJwtCookie(c, token)
	return utils.CreateEchoResponse(c, http.StatusCreated, response.AuthResponse{
		Token: token,
		Role:  role,
	})
}

func NewHandler(service account.Services) *Handler {
	return &Handler{service}
}
