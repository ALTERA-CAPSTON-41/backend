package handlers

import (
	"clinic-api/src/models"
	"clinic-api/src/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateDoctorHandler(c echo.Context) error {
	var doctorRequest models.DoctorRequest
	if err := c.Bind(&doctorRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	newDoctor := models.MapToNewDoctor(doctorRequest)
	if err := models.DB.Create(&newDoctor).Error; err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(
		c,
		http.StatusCreated,
		map[string]uuid.UUID{"id": newDoctor.UserID},
	)
}

func GetAllDoctorsHandler(c echo.Context) error {
	var doctors []models.Doctor
	name := c.QueryParam("name")
	nip := c.QueryParam("nip")

	query := "UPPER(name) LIKE '%" + strings.ToUpper(name) + "%'"
	if nip != "" {
		query = fmt.Sprintf("nip = '%s'", nip)
	}

	if err := models.DB.
		Preload("User").Preload("Polyclinic").
		Where(query).
		Find(&doctors).Error; err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(
		c,
		http.StatusOK,
		models.MapToDoctorBatchResponse(doctors),
	)
}

func GetDoctorByIDHandler(c echo.Context) error {
	id := c.Param("id")
	doctor := models.Doctor{}
	if err := models.DB.
		Preload("User").Preload("Polyclinic").
		Where("user_id", id).
		First(&doctor).Error; err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
		}

		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(
		c,
		http.StatusOK,
		models.MapToDoctorResponse(doctor),
	)
}

func EditDoctorByIDHandler(c echo.Context) error {
	id := c.Param("id")
	doctorRequest := models.DoctorRequest{}
	if err := c.Bind(&doctorRequest); err != nil {
		return utils.CreateEchoResponse(c, http.StatusBadRequest, nil)
	}

	doctor := models.MapToExistingDoctor(doctorRequest, id)
	editAction := models.DB.Where("user_id", id).Updates(&doctor)
	if editAction.RowsAffected < 1 {
		return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
	}

	err := editAction.Error
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}

func DeleteDoctorByIDHandler(c echo.Context) error {
	id := c.Param("id")
	deleteUser := models.DB.Where("id", id).Delete(new(models.User))
	if deleteUser.RowsAffected < 1 {
		return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
	}

	err := deleteUser.Error
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	deleteDoctor := models.DB.Where("user_id", id).Delete(new(models.Doctor))
	if deleteDoctor.RowsAffected < 1 {
		return utils.CreateEchoResponse(c, http.StatusNotFound, nil)
	}

	err = deleteDoctor.Error
	if err != nil {
		return utils.CreateEchoResponse(c, http.StatusInternalServerError, nil)
	}

	return utils.CreateEchoResponse(c, http.StatusNoContent, nil)
}
