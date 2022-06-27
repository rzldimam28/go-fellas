package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentResponse struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId  primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	BlogId  primitive.ObjectID `json:"blog_id,omitempty" bson:"blog_id,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
}