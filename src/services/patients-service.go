package PatientsService

import (
	"fmt"

	"github.com/luancgs/doctor-assistant-app/database/models"
	PatientsRepository "github.com/luancgs/doctor-assistant-app/src/repositories"
	"github.com/luancgs/doctor-assistant-app/src/utils"
)

func GetAllPatients() ([]models.Patient, error) {
	result, err := PatientsRepository.SelectAll()

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error getting all patients")
	}

	return result, nil
}

func GetPatientById(patientId string) (models.Patient, error) {
	id, err := utils.ConvertId(patientId)

	if err != nil {
		fmt.Println(err)
		return models.Patient{}, fmt.Errorf("error converting id to uint")
	}

	result, err := PatientsRepository.SelectById(id)

	if err != nil {
		fmt.Println(err)
		return models.Patient{}, fmt.Errorf("error getting patient")
	}

	return result, nil
}

func CreatePatient(patient models.Patient) (bool, error) {
	result, err := PatientsRepository.Insert(patient)

	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("error creating patient")
	}

	return result, nil
}

func UpdatePatient(patientId string, updatePatient models.Patient) (bool, error) {
	id, err := utils.ConvertId(patientId)

	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("error converting id to uint")
	}

	updatePatient.ID = id
	result, err := PatientsRepository.Update(updatePatient)

	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("error updating patient")
	}

	return result, nil
}

func DeletePatient(patientId string) (bool, error) {
	id, err := utils.ConvertId(patientId)

	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("error converting id to uint")
	}

	result, err := PatientsRepository.Delete(id)

	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("error deleting patient")
	}

	return result, nil
}
