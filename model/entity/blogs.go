package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blogs struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title string `json:"title,omitempty" bson:"title,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
	Comment *[]Comments `json:"comment,omitempty" bson:"comment,omitempty"`
	IsCom bool `json:"is_com,omitempty" bson:"is_com,omitempty"`
	LikedBy []primitive.ObjectID `json:"liked_by,omitempty" bson:"liked_by,omitempty"`
	LikedCount int `json:"liked_count,omitempty" bson:"liked_count,omitempty"`
}