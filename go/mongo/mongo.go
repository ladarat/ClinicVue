package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MgoDatabase struct
type MgoDatabase struct {
	MgoDatabase *mongo.Database
	Context     context.Context
	DBName      string
	client      *mongo.Client
}

const (
	// DBName is constant
	DBName = "clinic"
	// URI is constant
	URI = "mongodb+srv://admin:clinic@cluster0-vqe7a.gcp.mongodb.net/test?retryWrites=true&w=majority"
)

//NewDatabase instanet
func NewDatabase() (*MgoDatabase, error) {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Minute)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(DBName)

	return &MgoDatabase{db, ctx, DBName, client}, err
}
