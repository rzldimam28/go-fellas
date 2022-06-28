package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserResponse struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty"`
}