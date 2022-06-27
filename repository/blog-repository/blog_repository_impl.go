package blogrepository

import (
	"context"
	"errors"

	"github.com/rzldimam28/wlb-test/model/entity"
	"github.com/rzldimam28/wlb-test/model/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepositoryImpl struct {
	DB *mongo.Database
}

func NewBlogRepository(DB *mongo.Database) BlogRepository {
	return &BlogRepositoryImpl{
		DB: DB,
	}
}

func (blogRepository *BlogRepositoryImpl) FindAll(ctx context.Context) []entity.Blogs {
	cur, err := blogRepository.DB.Collection("blogs").Find(ctx, bson.M{})
	helper.PanicIfError(err)
	defer cur.Close(ctx)

	var blogs []entity.Blogs
	for cur.Next(ctx) {
		var blog entity.Blogs
		err := cur.Decode(&blog)
		helper.PanicIfError(err)
		blogs = append(blogs, blog)
	}

	return blogs
}

func (blogRepository *BlogRepositoryImpl) FindById(ctx context.Context, blogId primitive.ObjectID) (entity.Blogs, error) {
	blogToFind := bson.D{{Key: "_id", Value: blogId}}
	cur := blogRepository.DB.Collection("blogs").FindOne(ctx, blogToFind)
	
	var blog entity.Blogs
	err := cur.Decode(&blog)
	if err != nil {
		return blog, errors.New("can not find blog by id")
	}
	return blog, nil
}

func (blogRepository *BlogRepositoryImpl) Create(ctx context.Context, blog entity.Blogs) entity.Blogs {
	res, err := blogRepository.DB.Collection("blogs").InsertOne(ctx, blog)
	helper.PanicIfError(err)

	id := res.InsertedID
	oid := id.(primitive.ObjectID)
	blog.Id = oid
	return blog
}

func (blogRepository *BlogRepositoryImpl) Update(ctx context.Context, blog entity.Blogs) entity.Blogs {
	filter := bson.D{{Key: "_id", Value: blog.Id}}
	updatedBlog := bson.D{{Key: "$set", Value: bson.D{
		{Key: "title", Value: blog.Title},
		{Key: "content", Value: blog.Content},
	}}}
	_, err := blogRepository.DB.Collection("blogs").UpdateOne(ctx, filter, updatedBlog)
	helper.PanicIfError(err)

	return blog
}

func (blogRepository *BlogRepositoryImpl) Delete(ctx context.Context, blog entity.Blogs) {
	filter := bson.D{{Key: "_id", Value: blog.Id}}
	_, err := blogRepository.DB.Collection("blogs").DeleteOne(ctx, filter)
	helper.PanicIfError(err)
}

func (blogRepository *BlogRepositoryImpl) Like(ctx context.Context, blog entity.Blogs) entity.Blogs {
	filter := bson.D{{Key: "_id", Value: blog.Id}}
	updatedBlog := bson.D{{Key: "$set", Value: bson.D{
		{Key: "liked_count", Value: blog.LikedCount},
		{Key: "liked_by", Value: blog.LikedBy},
	}}}
	_, err := blogRepository.DB.Collection("blogs").UpdateOne(ctx, filter, updatedBlog)
	helper.PanicIfError(err)

	return blog
}

func (blogRepository *BlogRepositoryImpl) FindByTitle(ctx context.Context, title string, orderBy string, order int) ([]entity.Blogs, error) {
	blogToFind := bson.D{{Key: "title", Value: primitive.Regex{Pattern: title, Options: ""}}}
	// blogToFind := bson.D{{Key: "title", Value: title}}
	if orderBy == "" {
		orderBy = "_id"
	}
	if order == 0 {
		order = 1
	}
	opts := options.Find().SetSort(bson.D{{Key: orderBy, Value: order}})

	cur, err := blogRepository.DB.Collection("blogs").Find(ctx, blogToFind, opts)
	helper.PanicIfError(err)
	defer cur.Close(ctx)

	var blogs []entity.Blogs
	for cur.Next(ctx) {
		var blog entity.Blogs
		err := cur.Decode(&blog)
		if err != nil {
			return blogs, errors.New("can not find blogs by title")
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}