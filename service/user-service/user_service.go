package userservice

import (
	"context"

	"github.com/rzldimam28/wlb-test/model/web/request"
	"github.com/rzldimam28/wlb-test/model/web/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	FindAll(ctx context.Context) []response.UserResponse
	FindById(ctx context.Context, userId primitive.ObjectID) response.UserResponse
	FindByUsername(ctx context.Context, username string) response.UserResponse
	Create(ctx context.Context, request request.CreateUserRequest) response.UserResponse
	Update(ctx context.Context, request request.UpdateUserRequest) response.UserResponse
	Verify(ctx context.Context, userId primitive.ObjectID) response.UserResponse
	Delete(ctx context.Context, userId primitive.ObjectID)
}