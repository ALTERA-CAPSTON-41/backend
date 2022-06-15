package doctor_handlers

import (
	"clinic-api/src/app/doctor"
	"clinic-api/src/app/doctor/handlers/request"
	"clinic-api/src/app/doctor/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services doctor.Services
}

// onCreate
func (h *Handler) CreateDoctorHandler(c echo.Context) error {
	var doctorRequest request.NewRequest

	if err := c.Bind(&doctorRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	id, err := h.services.CreateDoctor(doctorRequest.MapToDomain())
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusCreated, response.CreateResponse{ID: id})
}

// onShowAll
func (h *Handler) ShowAllDoctorsHandler(c echo.Context) error {
	data, err := h.services.GetAllDoctors()

	if err != nil {
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

	if err := h.services.AmendDoctorByID(id, doctorRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
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
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service doctor.Services) *Handler {
	return &Handler{service}
}
