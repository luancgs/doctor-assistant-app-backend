package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID            uint
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	MedicalCode   sql.NullString `json:"medical_code"`
	Appointments  []Appointment  `json:"appointments"`
	Exams         []Exam         `json:"exams"`
	Prescriptions []Prescription `json:"prescriptions"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Patient struct {
	ID            uint
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	Age           uint8          `json:"age"`
	Birthday      time.Time      `json:"birthday"`
	Sex           string         `json:"sex"`
	Appointments  []Appointment  `json:"appointments"`
	Exams         []Exam         `json:"exams"`
	Prescriptions []Prescription `json:"prescriptions"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Appointment struct {
	ID        uint
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	Status    uint      `json:"status"`
	PatientID uint
	DoctorID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Exam struct {
	ID        uint
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	Status    uint      `json:"status"`
	Result    string    `json:"result"`
	PatientID uint
	DoctorID  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Prescription struct {
	ID              uint
	Title           string    `json:"title"`
	Expiration_date time.Time `json:"expiration_date"`
	PatientID       uint
	DoctorID        uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
