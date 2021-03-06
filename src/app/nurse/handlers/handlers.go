package nurse_handlers

import (
	"clinic-api/src/app/nurse"
	"clinic-api/src/app/nurse/handlers/request"
	"clinic-api/src/app/nurse/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services  nurse.Services
	validator *validator.Validate
}

// onCreate
func (h *Handler) CreateNurseHandler(c echo.Context) error {
	var nurseRequest request.NewRequest
	if err := c.Bind(&nurseRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if err := h.validator.Struct(nurseRequest); err != nil {
		var reason []string
		if strings.Contains(err.Error(), "email") {
			reason = append(reason, "email is invalid")
		}

		if strings.Contains(err.Error(), "Password") {
			reason = append(reason, "password must have at least 8 characters")
		}

		if !utils.ValidateName(nurseRequest.Name) {
			reason = append(reason, "name is invalid")
		}

		return utils.CreateEchoResponse(
			c,
			http.StatusBadRequest,
			response.ErrorResponse{Reason: reason},
		)
	}

	id, err := h.services.CreateNurse(nurseRequest.MapToDomain())
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

// onShowALl
func (h *Handler) ShowAllNursesHandler(c echo.Context) error {
	var polyclinic int
	page, _ := strconv.Atoi(c.QueryParam("page"))
	strPolyclinic := c.QueryParam("polyclinic")
	if strPolyclinic != "" {
		var err error
		polyclinic, err = strconv.Atoi(strPolyclinic)
		if err != nil {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
	}

	data, err := h.services.GetAllNurses(polyclinic, page)
	if err != nil {
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToBatchResponse(data))
}

// onShowOne
func (h *Handler) ShowNurseByIDHandler(c echo.Context) error {
	data, err := h.services.GetNurseByID(c.Param("id"))
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
func (h *Handler) AmendNurseByIDHandler(c echo.Context) error {
	id := c.Param("id")
	nurseRequest := request.UpdateRequest{}
	if err := c.Bind(&nurseRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if !utils.ValidateName(nurseRequest.Name) {
		return utils.CreateEchoResponse(
			c,
			http.StatusBadRequest,
			response.ErrorResponse{Reason: "name is invalid"},
		)
	}

	if err := h.services.AmendNurseByID(id, nurseRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

// onDelete
func (h *Handler) RemoveNurseByIDHandler(c echo.Context) error {
	if err := h.services.RemoveNurseByID(c.Param("id")); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service nurse.Services) *Handler {
	return &Handler{
		service,
		validator.New(),
	}
}
