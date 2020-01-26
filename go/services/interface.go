package services

//Patient struct
type Patient struct {
	ID                uint
	Firstname         string
	Lastname          string
	Nickname          string
	Sex               string
	Career            string
	PhoneNumber       string
	CurrentAddress    string
	WorkAddress       string
	RequiredDocuments string
	congenitalDisease string
	DrugAllergy       string
	EmergencyContact  string
	Relationship      string
}

//PatientService is interface
type PatientService interface {
	Create(version *Patient) error
}
