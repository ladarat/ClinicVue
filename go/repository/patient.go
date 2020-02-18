package repository

import (
	"time"

	models "github.com/ladarat/ClinicVue/go/models"
	mongo "github.com/ladarat/ClinicVue/go/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type patientRepository struct {
	db *mongo.MgoDatabase
}

const patientCollection = "patient"

// NewMongoPatientRepository is repository
func NewMongoPatientRepository(db *mongo.MgoDatabase) PatientRepository {
	return &patientRepository{
		db: db,
	}
}

func (pr *patientRepository) Create(patient *models.Patient) (*models.Patient, error) {
	coll := pr.db.MgoDatabase.Collection(patientCollection)
	patient.ID = primitive.NewObjectID()
	patient.CreatedAt = time.Now()
	patient.UpdatedAt = time.Now()
	_, err := coll.InsertOne(pr.db.Context, patient)
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func (pr *patientRepository) GetAll() ([]models.Patient, error) {
	coll := pr.db.MgoDatabase.Collection(patientCollection)
	patientList := []models.Patient{}
	p := models.Patient{}
	cursor, err := coll.Find(pr.db.Context, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(pr.db.Context) {
		cursor.Decode(&p)
		patientList = append(patientList, p)
	}

	return patientList, nil
}

func (pr *patientRepository) GetByID(id string) (*models.Patient, error) {
	coll := pr.db.MgoDatabase.Collection(patientCollection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	findResult := coll.FindOne(pr.db.Context, bson.M{"_id": objID})
	if err := findResult.Err(); err != nil {
		return nil, err
	}

	p := models.Patient{}
	err = findResult.Decode(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (pr *patientRepository) Delete(id string) error {
	coll := pr.db.MgoDatabase.Collection(patientCollection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = coll.DeleteOne(pr.db.Context, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}

func (pr *patientRepository) Update(patient *models.Patient, id string) (*models.Patient, error) {
	coll := pr.db.MgoDatabase.Collection(patientCollection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	patient.UpdatedAt = time.Now()
	_, err = coll.UpdateOne(pr.db.Context, bson.M{"_id": objID}, bson.M{
		"$set": patient,
	})
	if err != nil {
		return nil, err
	}

	return patient, nil
}
