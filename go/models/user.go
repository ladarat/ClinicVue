package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}
