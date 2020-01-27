package services

import (
	mongo "github.com/ladarat/ClinicVue/go/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type mongoPatientService struct {
	md *mongo.MgoDatabase
}

const patientCollection = "patient"

// NewPatientService will create and return
func NewPatientService(md *mongo.MgoDatabase) PatientService {
	return &mongoPatientService{
		md: md,
	}
}

func (mp *mongoPatientService) Create(patient *Patient) (*Patient, error) {
	coll := mp.md.MgoDatabase.Collection(patientCollection)
	patient.ID = primitive.NewObjectID()
	_, err := coll.InsertOne(mp.md.Context, patient)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (mp *mongoPatientService) GetAll() ([]Patient, error) {
	coll := mp.md.MgoDatabase.Collection(patientCollection)
	patientList := []Patient{}
	p := Patient{}
	cursor, err := coll.Find(mp.md.Context, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(mp.md.Context) {
		cursor.Decode(&p)
		patientList = append(patientList, p)
	}

	return patientList, nil
}

func (mp *mongoPatientService) GetByID(id string) (*Patient, error) {
	coll := mp.md.MgoDatabase.Collection(patientCollection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	findResult := coll.FindOne(mp.md.Context, bson.M{"_id": objID})
	if err := findResult.Err(); err != nil {
		return nil, err
	}

	p := Patient{}
	err = findResult.Decode(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (mp *mongoPatientService) Delete(id string) error {
	coll := mp.md.MgoDatabase.Collection(patientCollection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = coll.DeleteOne(mp.md.Context, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}

func (mp *mongoPatientService) Update(patient *Patient) error {

	return nil
}
