package services

import "go.mongodb.org/mongo-driver/bson/primitive"

//Patient struct
type Patient struct {
	ID                primitive.ObjectID `bson:"_id"`
	Firstname         string             `bson:"firstname"`
	Lastname          string             `bson:"lastname"`
	Nickname          string             `bson:"nickname"`
	Sex               string             `bson:"sex"`
	Career            string             `bson:"career"`
	PhoneNumber       string             `bson:"phone_number"`
	CurrentAddress    string             `bson:"current_address"`
	WorkAddress       string             `bson:"work_address"`
	RequiredDocuments string             `bson:"required_documents"`
	CongenitalDisease string             `bson:"congenital_disease"`
	DrugAllergy       string             `bson:"drug_allergy"`
	EmergencyContact  string             `bson:"emergency_contact"`
	Relationship      string             `bson:"relationship"`
}

//PatientService is interface
type PatientService interface {
	Create(patient *Patient) (*Patient, error)
	GetAll() ([]Patient, error)
	GetByID(id string) (*Patient, error)
	Delete(id string) error
	Update(patient *Patient) error
}
