package request

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateCommentRequest struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	UserId   primitive.ObjectID `json:"user_id" bson:"user_id"`
	BlogId   primitive.ObjectID `json:"blog_id" bson:"blog_id"`
	Content  string             `json:"content" bson:"content"`
}