package userrepository

import (
	"context"

	"github.com/rzldimam28/wlb-test/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	FindAll(ctx context.Context) []entity.Users
	FindById(ctx context.Context, userId primitive.ObjectID) (entity.Users, error)
	FindByUsername(ctx context.Context, username string) entity.Users
	Create(ctx context.Context, user entity.Users) entity.Users
	Update(ctx context.Context, user entity.Users) entity.Users
	Verify(ctx context.Context, user entity.Users) entity.Users
	Delete(ctx context.Context, user entity.Users)
}