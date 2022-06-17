package nurse_handlers

import (
	"clinic-api/src/app/nurse"
	"clinic-api/src/app/nurse/handlers/request"
	"clinic-api/src/app/nurse/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services nurse.Services
}

// onCreate
func (h *Handler) CreateNurseHandler(c echo.Context) error {
	var nurseRequest request.NewRequest
	if err := c.Bind(&nurseRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	id, err := h.services.CreateNurse(nurseRequest.MapToDomain())
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusCreated, response.CreateResponse{ID: id})
}

// onShowALl
func (h *Handler) ShowAllNursesHandler(c echo.Context) error {
	data, err := h.services.GetAllNurses()
	if err != nil {
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

	if err := h.services.AmendNurseByID(id, nurseRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
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
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service nurse.Services) *Handler {
	return &Handler{service}
}
