package doctor_handlers

import (
	"clinic-api/src/app/doctor"
	"clinic-api/src/app/doctor/handlers/request"
	"clinic-api/src/app/doctor/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services  doctor.Services
	validator *validator.Validate
}

// onCreate
func (h *Handler) CreateDoctorHandler(c echo.Context) error {
	var doctorRequest request.NewRequest

	if err := c.Bind(&doctorRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if err := h.validator.Struct(doctorRequest); err != nil {
		var reason []string
		if strings.Contains(err.Error(), "email") {
			reason = append(reason, "email is invalid")
		}

		if strings.Contains(err.Error(), "Password") {
			reason = append(reason, "password must have at least 8 characters")
		}

		if !utils.ValidateName(doctorRequest.Name) {
			reason = append(reason, "name is invalid")
		}

		return utils.CreateEchoResponse(
			c,
			http.StatusBadRequest,
			response.ErrorResponse{Reason: reason},
		)
	}

	id, err := h.services.CreateDoctor(doctorRequest.MapToDomain())
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
func (h *Handler) ShowAllDoctorsHandler(c echo.Context) error {
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

	data, err := h.services.GetAllDoctors(polyclinic, page)

	if err != nil {
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}
	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToBatchResponse(data))
}

// onShowOne
func (h *Handler) ShowDoctorByIDHandler(c echo.Context) error {
	data, err := h.services.GetDoctorByID(c.Param("id"))

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
func (h *Handler) AmendDoctorByIDHandler(c echo.Context) error {
	var doctorRequest request.UpdateRequest
	id := c.Param("id")

	if err := c.Bind(&doctorRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if !utils.ValidateName(doctorRequest.Name) {
		return utils.CreateEchoResponse(
			c,
			http.StatusBadRequest,
			response.ErrorResponse{Reason: "name is invalid"},
		)
	}

	if err := h.services.AmendDoctorByID(id, doctorRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

// onDelete
func (h *Handler) RemoveDoctorByIDHandler(c echo.Context) error {
	id := c.Param("id")

	if err := h.services.RemoveDoctorByID(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service doctor.Services) *Handler {
	return &Handler{
		service,
		validator.New(),
	}
}
