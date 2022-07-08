package prescription_handlers

import (
	"clinic-api/src/app/prescription"
	"clinic-api/src/app/prescription/handlers/request"
	"clinic-api/src/app/prescription/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services prescription.Services
}

// onCreate
func (h *Handler) CreatePrescriptionHandler(c echo.Context) error {
	var prescriptionRequest request.NewRequest
	if err := c.Bind(&prescriptionRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	id, err := h.services.CreatePrescription(prescriptionRequest.MapToDomain())
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

// onShowAllByMedicalRecordID
func (h *Handler) ShowAllPresciptionsByIDHandler(c echo.Context) error {
	data, err := h.services.FindPrescriptionsByID(c.Param("id"))
	if err != nil {
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToBatchResponse(data))
}

// onUpdate
func (h *Handler) AmendPrescriptionByIDHandler(c echo.Context) error {
	id := c.Param("id")
	prescriptionRequest := request.UpdateRequest{}
	if err := c.Bind(&prescriptionRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if err := h.services.
		AmendPrescriptionByID(id, prescriptionRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

// onDelete
func (h *Handler) RemovePrescriptionByIDHandler(c echo.Context) error {
	if err := h.services.RemovePrescriptionByID(c.Param("id")); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service prescription.Services) *Handler {
	return &Handler{service}
}
