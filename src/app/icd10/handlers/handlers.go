package icd10_handlers

import (
	"clinic-api/src/app/icd10"
	"clinic-api/src/app/icd10/handlers/response"
	"clinic-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services icd10.Services
}

// OnShowAll
func (h *Handler) FindICD10ByCodeHandler(c echo.Context) error {
	code := c.Param("code")
	if len(code) < 2 {
		return utils.CreateEchoResponse(
			c,
			http.StatusBadRequest,
			response.ErrorResponse{Reason: "must provide at least 2 characters"},
		)
	}

	data, err := h.services.FindICD10ByCode(code)
	if err != nil {
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToBatchResponse(data))
}

func NewHandler(service icd10.Services) *Handler {
	return &Handler{service}
}
