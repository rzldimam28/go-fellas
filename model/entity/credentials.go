package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Credential struct {
	Id primitive.ObjectID `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}