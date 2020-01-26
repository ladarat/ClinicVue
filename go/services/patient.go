package services

type mongoPatientService struct{}

// NewPatientService will create and return
func NewPatientService() PatientService {
	return &mongoPatientService{}
}

func (cs *mongoPatientService) Create(version *Patient) error {

	return nil
}
