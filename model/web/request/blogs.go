package request

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateBlogRequest struct {
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	IsCom bool `json:"is_com" bson:"is_com"`
}

type UpdateBlogRequest struct {
	Id primitive.ObjectID `json:"_id" bson:"_id"`
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	IsCom bool `json:"is_com" bson:"is_com"`
	LikedBy []primitive.ObjectID `json:"liked_by" bson:"liked_by"`
	LikedCount int `json:"liked_count" bson:"liked_count"`
}