package repository

import (
	models "github.com/ladarat/ClinicVue/go/models"
)

//UserRepository is interface
type UserRepository interface {
	GetAll() ([]models.User, error)
}

//PatientRepository is interface
type PatientRepository interface {
	Create(patient *models.Patient) (*models.Patient, error)
	GetAll() ([]models.Patient, error)
	GetByID(id string) (*models.Patient, error)
	Delete(id string) error
	Update(patient *models.Patient, id string) (*models.Patient, error)
}
