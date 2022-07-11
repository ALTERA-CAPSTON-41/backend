package medicalrecord_handlers

import (
	medicalrecord "clinic-api/src/app/medical_record"
	"clinic-api/src/app/medical_record/handlers/request"
	"clinic-api/src/app/medical_record/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services medicalrecord.Services
}

// onAmend
func (h *Handler) AmendMedicalRecordByIDHandler(c echo.Context) error {
	var medicalrecordRequest request.Request

	token := utils.GetJwtTokenFromRequest(c)
	claims, _ := utils.ExtractClaims((token))
	id := c.Param("id")

	if err := c.Bind(&medicalrecordRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	domain := medicalrecordRequest.MapToDomain()
	domain.Doctor.ID = uuid.MustParse(claims.Id)

	if err := h.services.AmendMedicalRecordByID(domain, id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}

		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, err.Error())
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

// onCreate
func (h *Handler) CreateMedicalRecordHandler(c echo.Context) error {
	var medicalrecordRequest request.Request

	token := utils.GetJwtTokenFromRequest(c)
	claims, _ := utils.ExtractClaims(token)

	if err := c.Bind(&medicalrecordRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	domain := medicalrecordRequest.MapToDomain()
	domain.Doctor.ID = uuid.MustParse(claims.Id)

	id, err := h.services.CreateMedicalRecord(domain)
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

// onShowByPatientID
func (h *Handler) ShowMedicalRecordByPatientIDHandler(c echo.Context) error {
	data, err := h.services.FindMedicalRecordByPatientID(c.Param("id"))

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}

		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToBatchResponse(data))
}

// onShowByPateintNIK
func (h *Handler) ShowMedicalRecordByPatientNIKHandler(c echo.Context) error {
	data, err := h.services.FindMedicalRecordByPatientNIK(c.Param("nik"))

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}

		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToBatchResponse(data))
}

// onShowByID
func (h *Handler) ShowMedicalRecordByIDHandler(c echo.Context) error {
	data, err := h.services.FindMedicalRecordByID(c.Param("id"))

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}

		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToResponse(*data))
}

// onRemove
func (h *Handler) RemoveMedicalRecordByIDHandler(c echo.Context) error {
	id := c.Param("id")

	if err := h.services.RemoveMedicalRecordByID(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}

		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service medicalrecord.Services) *Handler {
	return &Handler{service}
}
