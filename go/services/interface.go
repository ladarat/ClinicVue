package services

import (
	models "github.com/ladarat/ClinicVue/go/models"
)

//UserService is interface
type UserService interface {
	GetAll() ([]models.User, error)
}

//PatientService is interface
type PatientService interface {
	Create(patient *models.Patient) (*models.Patient, error)
	GetAll() ([]models.Patient, error)
	GetByID(id string) (*models.Patient, error)
	Delete(id string) error
	Update(patient *models.Patient, id string) (*models.Patient, error)
}
