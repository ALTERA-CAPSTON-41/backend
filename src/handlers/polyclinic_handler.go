package handlers

import (
	"clinic-api/src/models"
	"clinic-api/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreatePolyclinicHandler(c echo.Context) error {
	var polyclinicRequest models.PolyclinicRequest
	if err := c.Bind(&polyclinicRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	newPolyclinic := models.MapToNewPolyclinic(polyclinicRequest)
	if err := models.DB.Create(&newPolyclinic).Error; err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusCreated, map[string]int{"id": newPolyclinic.ID})
}

func GetAllPolyclinicHandler(c echo.Context) error {
	var polyclinics []models.Polyclinic
	if err := models.DB.Find(&polyclinics).Error; err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, models.MapToPolyclinicBatch(polyclinics))
}

func GetPolyclinicByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	polyclinic := models.Polyclinic{}
	if err := models.DB.First(&polyclinic, id).Error; err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}

		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusOK, models.MapToPolyclinic(polyclinic))
}

func EditPolyclinicByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	polyclinicRequest := models.PolyclinicRequest{}
	if err := c.Bind(&polyclinicRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	polyclinic := models.MapToExistingPolyclinic(polyclinicRequest, id)
	editAction := models.DB.Where("id", id).Updates(&polyclinic)
	if editAction.RowsAffected < 1 {
		return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
	}

	err := editAction.Error
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func DeletePolyclinicByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	polyclinic := models.Polyclinic{}
	deleteAction := models.DB.Where("id", id).Delete(&polyclinic)
	if deleteAction.RowsAffected < 1 {
		return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
	}

	err := deleteAction.Error
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}