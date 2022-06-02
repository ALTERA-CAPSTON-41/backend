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

func CreatePatientHandler(c echo.Context) error {
	status := http.StatusCreated
	patientRequest := models.PatientRequest{}
	if err := c.Bind(&patientRequest); err != nil {
		status = http.StatusBadRequest
		return utils.CreateEchoResponse(c, status, nil)
	}

	newPatient := models.MapToNewPatient(patientRequest)
	if err := models.DB.Create(&newPatient).Error; err != nil {
		status = http.StatusInternalServerError
		return utils.CreateEchoResponse(c, status, nil)
	}

	return utils.CreateEchoResponse(c, status, map[string]uuid.UUID{"id": newPatient.ID})
}

func GetAllPatientsHandler(c echo.Context) error {
	status := http.StatusOK
	name := c.QueryParam("name")
	nik := c.QueryParam("nik")
	patients := []models.Patient{}

	query := "UPPER(name) LIKE '%" + strings.ToUpper(name) + "%'"
	if nik != "" {
		query = fmt.Sprint("nik = ", nik)
	}

	if err := models.DB.Where(query).Find(&patients).Error; err != nil {
		status = http.StatusInternalServerError
		return utils.CreateEchoResponse(c, status, nil)
	}

	return utils.CreateEchoResponse(c, status, models.MapToPatientBatch(patients))
}

func GetPatientByIDHandler(c echo.Context) error {
	status := http.StatusOK
	id := c.Param("id")
	patient := models.Patient{}
	if err := models.DB.Where("id", id).First(&patient).Error; err != nil {
		if strings.Contains(err.Error(), "not found") {
			status = http.StatusNotFound
			return utils.CreateEchoResponse(c, status, nil)
		}

		status = http.StatusInternalServerError
		return utils.CreateEchoResponse(c, status, nil)
	}

	return utils.CreateEchoResponse(c, status, models.MapToPatient(patient))
}

func EditDoctorByIDHandler(c echo.Context) error {
	status := http.StatusNoContent
	id := c.Param("id")
	patientRequest := models.PatientRequest{}
	if err := c.Bind(&patientRequest); err != nil {
		status = http.StatusBadRequest
		return utils.CreateEchoResponse(c, status, nil)
	}

	patient := models.MapToExistingPatient(patientRequest, id)
	editAction := models.DB.Where("id", id).Updates(&patient)
	if editAction.RowsAffected < 1 {
		status = http.StatusNotFound
		return utils.CreateEchoResponse(c, status, nil)
	}

	err := editAction.Error
	if err != nil {
		status = http.StatusInternalServerError
		return utils.CreateEchoResponse(c, status, nil)
	}

	return utils.CreateEchoResponse(c, status, nil)
}

func DeletePatientByIDHandler(c echo.Context) error {
	status := http.StatusNoContent
	id := c.Param("id")
	patient := models.Patient{}
	deleteAction := models.DB.Where("id", id).Delete(&patient)
	if deleteAction.RowsAffected < 1 {
		status = http.StatusNotFound
		return utils.CreateEchoResponse(c, status, nil)
	}

	err := deleteAction.Error
	if err != nil {
		status = http.StatusInternalServerError
		return utils.CreateEchoResponse(c, status, nil)
	}

	return utils.CreateEchoResponse(c, status, nil)
}
