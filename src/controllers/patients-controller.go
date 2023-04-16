package PatientsController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luancgs/doctor-assistant-app/database/models"
	PatientsService "github.com/luancgs/doctor-assistant-app/src/services"
)

func GetAllPatients(c *gin.Context) {
	output, err := PatientsService.GetAllPatients()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, output)
}

func GetPatientById(c *gin.Context) {
	output, err := PatientsService.GetPatientById(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, output)
}

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	c.Bind(&patient)

	output, err := PatientsService.CreatePatient(patient)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, output)
}

func UpdatePatient(c *gin.Context) {
	var patient models.Patient
	c.Bind(&patient)

	output, err := PatientsService.UpdatePatient(c.Param("id"), patient)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, output)
}

func DeletePatient(c *gin.Context) {
	output, err := PatientsService.DeletePatient(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, output)
}
