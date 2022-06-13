package handlers

import (
	"clinic-api/src/app/polyclinic"
	"clinic-api/src/app/polyclinic/handlers/request"
	"clinic-api/src/app/polyclinic/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services polyclinic.Services
}

// onCreate
func (h *Handler) CreatePolyclinicHandler(c echo.Context) error {
	var polyclinicRequest request.Request
	if err := c.Bind(&polyclinicRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	id, err := h.services.CreatePolyclinic(polyclinicRequest.MapToDomain())
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(
		c,
		http.StatusCreated,
		response.CreateResponse{ID: id},
	)
}

// onShowAll
func (h *Handler) ShowPolyclinicHandler(c echo.Context) error {
	data, err := h.services.GetAllPolyclinics()
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToBatchResponse(data))
}

// onShowOne
func (h *Handler) ShowPolyclinicByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := h.services.GetPolyclinicByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToResponse(*data))
}

// onUpdate
func (h *Handler) AmendPolyclinicByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	polyclinicRequest := request.Request{}
	if err := c.Bind(&polyclinicRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	err := h.services.AmendPolyclinicByID(id, polyclinicRequest.MapToDomain())
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

// onDelete
func (h *Handler) RemovePolyclinicByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.services.RemovePolyclinicByID(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service polyclinic.Services) *Handler {
	return &Handler{service}
}
