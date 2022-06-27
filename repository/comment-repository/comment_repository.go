package commentrepository

import (
	"context"

	"github.com/rzldimam28/wlb-test/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comments) entity.Comments
	FindByBlogId(ctx context.Context, blogId primitive.ObjectID) []entity.Comments
}