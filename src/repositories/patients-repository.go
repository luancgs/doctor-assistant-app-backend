package PatientsRepository

import (
	"github.com/luancgs/doctor-assistant-app/database"
	"github.com/luancgs/doctor-assistant-app/database/models"
)

func SelectAll() ([]models.Patient, error) {
	var db = database.GetInstance()
	var patients []models.Patient
	result := db.Find(&patients)

	if result.Error != nil {
		return nil, result.Error
	}

	return patients, nil
}

func SelectById(patientId uint) (models.Patient, error) {
	var db = database.GetInstance()
	var patient models.Patient
	result := db.Where("id = ?", patientId).First(&patient)

	if result.Error != nil {
		return models.Patient{}, result.Error
	}

	return patient, nil
}

func Insert(patient models.Patient) (bool, error) {
	var db = database.GetInstance()
	result := db.Create(&patient)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func Update(patient models.Patient) (bool, error) {
	var db = database.GetInstance()
	result := db.Updates(&patient)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func Delete(patientId uint) (bool, error) {
	var db = database.GetInstance()
	var patient models.Patient
	result := db.Where("id = ?", patientId).Delete(&patient)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
