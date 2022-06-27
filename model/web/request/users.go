package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateUserRequest struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Status   string `json:"status" bson:"status"`
}

type UpdateUserRequest struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Username    string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Status   string             `json:"status" bson:"status"`
}

type VerifyUser struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Status   string             `json:"status" bson:"status"`
}