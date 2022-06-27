package response

import (
	"github.com/rzldimam28/wlb-test/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogResponse struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
	Comments *[]entity.Comments `json:"comments,omitempty" bson:"comments,omitempty"`
	IsCom bool `json:"is_com,omitempty" bson:"is_com,omitempty"`
	LikedBy []primitive.ObjectID `json:"liked_by,omitempty" bson:"liked_by,omitempty"`
	LikedCount int `json:"liked_count,omitempty" bson:"liked_count,omitempty"`
}
