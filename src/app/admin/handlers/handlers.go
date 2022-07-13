package admin_handlers

import (
	"clinic-api/src/app/admin"
	"clinic-api/src/app/admin/handlers/request"
	"clinic-api/src/app/admin/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services  admin.Services
	validator *validator.Validate
}

// onCreate
func (h *Handler) CreateAdminHandler(c echo.Context) error {
	var adminRequest request.NewRequest

	if err := c.Bind(&adminRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if err := h.validator.Struct(adminRequest); err != nil {
		var reason []string
		if strings.Contains(err.Error(), "email") {
			reason = append(reason, "email is invalid")
		}

		if strings.Contains(err.Error(), "Password") {
			reason = append(reason, "password must have at least 8 characters")
		}

		if !utils.ValidateName(adminRequest.Name) {
			reason = append(reason, "name is invalid")
		}

		return utils.CreateEchoResponse(
			c,
			http.StatusBadRequest,
			response.ErrorResponse{Reason: reason},
		)
	}

	id, err := h.services.CreateAdmin(adminRequest.MapToDomain())
	if err != nil {
		if strings.Contains(err.Error(), "already used") {
			return utils.CreateEchoResponse(
				c,
				http.StatusBadRequest,
				response.ErrorResponse{Reason: err.Error()},
			)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusCreated, response.CreateResponse{ID: id})
}

// onShowAll
func (h *Handler) ShowAllAdminsHandler(c echo.Context) error {
	data, err := h.services.GetAllAdmins()

	if err != nil {
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}
	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToBatchResponse(data))
}

// onShowOne
func (h *Handler) ShowAdminByIDHandler(c echo.Context) error {
	data, err := h.services.GetAdminByID(c.Param("id"))

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToResponse(*data))
}

// onUpdate
func (h *Handler) AmendAdminByIDHandler(c echo.Context) error {
	var adminRequest request.UpdateRequest
	id := c.Param("id")

	if err := c.Bind(&adminRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if !utils.ValidateName(adminRequest.Name) {
		return utils.CreateEchoResponse(
			c,
			http.StatusBadRequest,
			response.ErrorResponse{Reason: "name is invalid"},
		)
	}

	if err := h.services.AmendAdminByID(id, adminRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

// onDelete
func (h *Handler) RemoveAdminByIDHandler(c echo.Context) error {
	id := c.Param("id")

	if err := h.services.RemoveAdminByID(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service admin.Services) *Handler {
	return &Handler{
		service,
		validator.New(),
	}
}
