package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luancgs/doctor-assistant-app/database"
	"github.com/luancgs/doctor-assistant-app/src/router"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	r := gin.Default()

	router.GetRoutes(r)

	err = r.Run() // listen and serve on 0.0.0.0:8080

	if err != nil {
		log.Fatal(err)
	}
}
