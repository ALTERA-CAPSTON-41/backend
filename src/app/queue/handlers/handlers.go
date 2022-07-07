package queue_handlers

import (
	"clinic-api/src/app/queue"
	"clinic-api/src/app/queue/handlers/request"
	"clinic-api/src/app/queue/handlers/response"
	"clinic-api/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services queue.Services
}

// onCreate
func (h *Handler) CreateQueueHandler(c echo.Context) error {
	var queueRequest request.NewRequest
	if err := c.Bind(&queueRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	id, err := h.services.CreateQueue(queueRequest.MapToDomain())
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

// onShowAll
func (h *Handler) ShowAllQueuesHandler(c echo.Context) error {
	polyclinic := c.QueryParam("polyclinic")
	fromDate := c.QueryParam("from-date")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	data, err := h.services.GetAllQueues(polyclinic, fromDate, page)
	if err != nil {
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, response.MapToBatchResponse(data))
}

// onUpdate
func (h *Handler) AmendQueueByIDHandler(c echo.Context) error {
	id := c.Param("id")
	queueRequest := request.UpdateRequest{}
	if err := c.Bind(&queueRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	if err := h.services.AmendQueueByID(id, queueRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

// onDelete
func (h *Handler) RemoveQueueByIDHandler(c echo.Context) error {
	if err := h.services.RemoveQueueByID(c.Param("id")); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}
		utils.CreateLog(c, err.Error())
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func NewHandler(service queue.Services) *Handler {
	return &Handler{service}
}
