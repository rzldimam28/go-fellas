package blogrepository

import (
	"context"

	"github.com/rzldimam28/wlb-test/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogRepository interface {
	FindAll(ctx context.Context) []entity.Blogs
	FindById(ctx context.Context, blogId primitive.ObjectID) (entity.Blogs, error)
	Create(ctx context.Context, blog entity.Blogs) entity.Blogs
	Update(ctx context.Context, blog entity.Blogs) entity.Blogs
	Delete(ctx context.Context, blog entity.Blogs)
	Like(ctx context.Context, blog entity.Blogs) entity.Blogs
	FindByTitle(ctx context.Context, title string, orderBy string, order int) ([]entity.Blogs, error)
}