package repository

import (
	models "github.com/ladarat/ClinicVue/go/models"
	mongo "github.com/ladarat/ClinicVue/go/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type userRepository struct {
	db *mongo.MgoDatabase
}

const userCollection = "user"

// NewMongoUserRepository is repository
func NewMongoUserRepository(db *mongo.MgoDatabase) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (pr *userRepository) GetAll() ([]models.User, error) {
	coll := pr.db.MgoDatabase.Collection(userCollection)
	patientList := []models.User{}
	p := models.User{}
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
