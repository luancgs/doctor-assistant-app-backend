package database

import (
	"fmt"
	"os"

	"github.com/luancgs/doctor-assistant-app/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"), os.Getenv("DB_TIMEZONE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Doctor{})
	db.AutoMigrate(&models.Patient{})
	db.AutoMigrate(&models.Appointment{})

	DB = db
}

func GetInstance() *gorm.DB {
	return DB
}
