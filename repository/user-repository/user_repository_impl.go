package userrepository

import (
	"context"
	"errors"

	"github.com/rzldimam28/wlb-test/model/entity"
	"github.com/rzldimam28/wlb-test/model/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	DB *mongo.Database
}

func NewUserRepository(DB *mongo.Database) UserRepository {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

func (userRepository *UserRepositoryImpl) FindAll(ctx context.Context) []entity.Users {
	cur, err := userRepository.DB.Collection("users").Find(ctx, bson.M{})
	helper.PanicIfError(err)
	defer cur.Close(ctx)

	var users []entity.Users
	for cur.Next(ctx) {
		var user entity.Users
		err := cur.Decode(&user)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}

func (userRepository *UserRepositoryImpl) FindById(ctx context.Context, userId primitive.ObjectID) (entity.Users, error) {
	userToFind := bson.D{{Key: "_id", Value: userId}}
	cur := userRepository.DB.Collection("users").FindOne(ctx, userToFind)
	
	var user entity.Users
	err := cur.Decode(&user)
	if err != nil {
		return user, errors.New("can not find user by id")
	}
	return user, nil
}

func (userRepository *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (entity.Users, error) {
	userToFind := bson.D{{Key: "username", Value: username}}
	cur := userRepository.DB.Collection("users").FindOne(ctx, userToFind)
	
	var user entity.Users
	err := cur.Decode(&user)
	if err != nil {
		return user, errors.New("can not find user by username")
	}
	return user, nil
}

func (userRepository *UserRepositoryImpl) Create(ctx context.Context, user entity.Users) entity.Users {
	res, err := userRepository.DB.Collection("users").InsertOne(ctx, user)
	helper.PanicIfError(err)

	id := res.InsertedID
	oid := id.(primitive.ObjectID)
	user.Id = oid
	return user
}

func (userRepository *UserRepositoryImpl) Update(ctx context.Context, user entity.Users) entity.Users {
	filter := bson.D{{Key: "_id", Value: user.Id}}
	updatedUser := bson.D{{Key: "$set", Value: bson.D{
		{Key: "username", Value: user.Username},
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
		{Key: "status", Value: user.Status},
	}}}
	_, err := userRepository.DB.Collection("users").UpdateOne(ctx, filter, updatedUser)
	helper.PanicIfError(err)

	return user
}

func (userRepository *UserRepositoryImpl) Verify(ctx context.Context, user entity.Users) entity.Users {
	filter := bson.D{{Key: "_id", Value: user.Id}}
	updatedUser := bson.D{{Key: "$set", Value: bson.D{
		{Key: "status", Value: "Aktif"},
	}}}
	_, err := userRepository.DB.Collection("users").UpdateOne(ctx, filter, updatedUser)
	helper.PanicIfError(err)

	return user
}

func (userRepository *UserRepositoryImpl) Delete(ctx context.Context, user entity.Users) {
	filter := bson.D{{Key: "_id", Value: user.Id}}
	_, err := userRepository.DB.Collection("users").DeleteOne(ctx, filter)
	helper.PanicIfError(err)
}