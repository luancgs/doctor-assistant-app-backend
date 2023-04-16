package router

import (
	"github.com/gin-gonic/gin"
	PatientsController "github.com/luancgs/doctor-assistant-app/src/controllers"
)

func GetRoutes(r *gin.Engine) {
	r.GET("/patients", PatientsController.GetAllPatients)
	r.GET("/patients/:id", PatientsController.GetPatientById)
	r.POST("/patients", PatientsController.CreatePatient)
	r.PUT("/patients/:id", PatientsController.UpdatePatient)
	r.DELETE("/patients/:id", PatientsController.DeletePatient)
}
