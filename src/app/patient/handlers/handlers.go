package patient_handlers

import (
	"clinic-api/src/app/patient"
	"clinic-api/src/app/patient/handlers/request"
	"clinic-api/src/app/patient/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services  patient.Services
	validator *validator.Validate
}

// onCreate
func (h *Handler) CreatePatientHandler(c echo.Context) error {
	var patientRequest request.Request

	if err := c.Bind(&patientRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if err := h.validator.Struct(patientRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	id, err := h.services.CreatePatient(patientRequest.MapToDomain())
	if err != nil {
		utils.CreateLog(c, err.Error())
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
	page, _ := strconv.Atoi(c.QueryParam("page"))
	var patientRequest request.Request
	patientRequest.Name = c.QueryParam("name")
	patientRequest.NIK = c.QueryParam("nik")

	result, err := h.services.HuntPatientByNameOrNIKOrAll(patientRequest.MapToDomain(), page)
	if err != nil {
		utils.CreateLog(c, err.Error())
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
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToResponse(*data))
}

// onUpdate
func (h *Handler) AmendPatientByIDHandler(c echo.Context) error {
	var patientRequest request.Request
	id := c.Param("id")

	if err := c.Bind(&patientRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if err := h.validator.Struct(patientRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if err := h.services.AmendPatientByID(id, patientRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
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
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service patient.Services) *Handler {
	return &Handler{
		service,
		validator.New(),
	}
}
