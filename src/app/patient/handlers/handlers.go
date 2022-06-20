package patient_handlers

import (
	"clinic-api/src/app/patient"
	"clinic-api/src/app/patient/handlers/request"
	"clinic-api/src/app/patient/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services patient.Services
}

// onCreate
func (h *Handler) CreatePatientHandler(c echo.Context) error {
	var patientRequest request.Request

	if err := c.Bind(&patientRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	id, err := h.services.CreatePatient(patientRequest.MapToDomain())
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(
		c,
		http.StatusCreated,
		response.CreateResponse{ID: id},
	)
}

// onHunt
func (h *Handler) HuntPatientByNameOrNIKOrAllHandler(c echo.Context) error {
	var patientRequest request.Request
	patientRequest.Name = c.QueryParam("name")
	patientRequest.NIK = c.QueryParam("nik")

	result, err := h.services.HuntPatientByNameOrNIKOrAll(patientRequest.MapToDomain())
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(
		c,
		http.StatusOK,
		response.MapToBatchResponse(result),
	)
}

// onShowOne
func (h *Handler) ShowPatientByIDHandler(c echo.Context) error {
	data, err := h.services.GetPatientByID(c.Param("id"))

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToResponse(*data))
}

// onUpdate
func (h *Handler) AmendPatientByIDHandler(c echo.Context) error {
	var adminRequest request.Request
	id := c.Param("id")

	if err := c.Bind(&adminRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if err := h.services.AmendPatientByID(id, adminRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

// onDelete
func (h *Handler) RemovePatientByIDHandler(c echo.Context) error {
	id := c.Param("id")

	if err := h.services.RemovePatientByID(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service patient.Services) *Handler {
	return &Handler{service}
}