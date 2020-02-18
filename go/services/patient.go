package services

import (
	models "github.com/ladarat/ClinicVue/go/models"
	repository "github.com/ladarat/ClinicVue/go/repository"
)

type patientService struct {
	patientRepo repository.PatientRepository
}

// NewPatientService is service
func NewPatientService(patientRepo repository.PatientRepository) PatientService {
	return &patientService{
		patientRepo: patientRepo,
	}
}

func (ps *patientService) Create(patient *models.Patient) (*models.Patient, error) {
	patient, err := ps.patientRepo.Create(patient)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (ps *patientService) GetAll() ([]models.Patient, error) {
	patients, err := ps.patientRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (ps *patientService) GetByID(id string) (*models.Patient, error) {
	patient, err := ps.patientRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (ps *patientService) Delete(id string) error {
	err := ps.patientRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (ps *patientService) Update(patient *models.Patient, id string) (*models.Patient, error) {
	patients, err := ps.patientRepo.Update(patient, id)
	if err != nil {
		return nil, err
	}

	return patients, nil
}
