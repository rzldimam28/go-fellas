package commentrepository

import (
	"context"

	"github.com/rzldimam28/wlb-test/model/entity"
	"github.com/rzldimam28/wlb-test/model/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRepositoryImpl struct {
	DB *mongo.Database
}

func NewCommentRepository(DB *mongo.Database) CommentRepository {
	return &CommentRepositoryImpl{
		DB: DB,
	}
}

func (commentRepository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comments) entity.Comments {
	res, err := commentRepository.DB.Collection("comments").InsertOne(ctx, comment)
	helper.PanicIfError(err)

	id := res.InsertedID
	oid := id.(primitive.ObjectID)
	comment.Id = oid
	return comment
}

func (commentRepository *CommentRepositoryImpl) FindByBlogId(ctx context.Context, blogId primitive.ObjectID) []entity.Comments {
	filter := bson.D{{Key: "blog_id", Value: blogId}}
	cur, err := commentRepository.DB.Collection("comments").Find(ctx, filter)
	helper.PanicIfError(err)
	defer cur.Close(ctx)

	var comments []entity.Comments
	for cur.Next(ctx) {
		var comment entity.Comments
		err := cur.Decode(&comment)
		helper.PanicIfError(err)
		comments = append(comments, comment)
	}

	return comments
}

