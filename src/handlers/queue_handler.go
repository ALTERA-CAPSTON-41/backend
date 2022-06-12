package handlers

import (
	"clinic-api/src/models"
	"clinic-api/src/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateQueueHandler(c echo.Context) error {
	var queueRequest models.QueueRequest
	if err := c.Bind(&queueRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	var queueNumber int
	if err := models.DB.Table("queues").Select("COALESCE(MAX(daily_queue_number), 0)").
		Where("daily_queue_date = CURDATE() AND polyclinic_id = ?", queueRequest.PolyclinicID).
		Find(&queueNumber).Error; err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}
	queueNumber++

	newQueue := models.MapToNewQueue(queueRequest, queueNumber)
	if err := models.DB.Create(&newQueue).Error; err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusCreated, newQueue.ID)
}

func GetAllQueuesHandler(c echo.Context) error {
	var (
		queues       []models.Queue
		byPolyclinic string
		polyclinic   = c.QueryParam("polyclinic")
		from         = c.QueryParam("from")
		byDate       = " AND daily_queue_date = CURDATE()"
	)

	if polyclinic != "" {
		byPolyclinic = fmt.Sprint(" AND polyclinic_id = ", polyclinic)
	}

	if from != "" {
		byDate = fmt.Sprint(" AND daily_queue_date >= ", from)
	}

	if err := models.DB.
		Preload("Patient").Preload("Polyclinic").
		Order("daily_queue_date DESC, daily_queue_number").
		Where("service_done_at IS NULL" + byPolyclinic + byDate).
		Find(&queues).Error; err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(
		c,
		http.StatusOK,
		models.MapToQueueBatchResponse(queues),
	)
}

func GetQueueByIDHandler(c echo.Context) error {
	id := c.Param("id")
	queue := models.Queue{}
	if err := models.DB.
		Preload("Patient").Preload("Polyclinic").
		Where("id", id).
		First(&queue).Error; err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}

		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, models.MapToQueueResponse(queue))
}

func EditQueueByIDHandler(c echo.Context) error {
	id := c.Param("id")
	queueRequest := models.QueueRequest{}
	if err := c.Bind(&queueRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	queue := models.MapToExistingQueue(queueRequest, id)
	editAction := models.DB.Where("id", id).Updates(&queue)
	if editAction.RowsAffected < 1 {
		return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
	}

	err := editAction.Error
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func DeleteQueueByIDHandler(c echo.Context) error {
	id := c.Param("id")
	deleteAction := models.DB.Where("id", id).Delete(new(models.Queue))
	if deleteAction.RowsAffected < 1 {
		return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
	}

	err := deleteAction.Error
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}
