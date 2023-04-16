package router

import (
	"github.com/gin-gonic/gin"
	"github.com/luancgs/doctor-assistant-app/src/controllers"
)

func GetRoutes(r *gin.Engine) {
	r.GET("/patients", controllers.GetAllPatients)
	r.GET("/patients/:id", controllers.GetPatientById)
	r.POST("/patients", controllers.CreatePatient)
	r.PUT("/patients/:id", controllers.UpdatePatient)
	r.DELETE("/patients/:id", controllers.DeletePatient)
}
