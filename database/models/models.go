package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID           uint
	Name         string
	Email        string
	Password     string
	MedicalCode  sql.NullString
	Appointments []Appointment
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Patient struct {
	ID           uint
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Age          uint8     `json:"age"`
	Birthday     time.Time `json:"birthday"`
	Sex          string    `json:"sex"`
	Appointments []Appointment
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Appointment struct {
	ID        uint
	Title     string
	Date      time.Time
	PatientID uint
	DoctorID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TODO: add exam, prescription, document and payment structures
