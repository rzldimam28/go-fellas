package blogservice

import (
	"context"

	"github.com/rzldimam28/wlb-test/model/web/request"
	"github.com/rzldimam28/wlb-test/model/web/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogService interface {
	FindAll(ctx context.Context) []response.BlogResponse
	FindById(ctx context.Context, blogId primitive.ObjectID) response.BlogResponse
	Create(ctx context.Context, request request.CreateBlogRequest, id primitive.ObjectID) response.BlogResponse
	Update(ctx context.Context, request request.UpdateBlogRequest) response.BlogResponse
	Delete(ctx context.Context, blogId primitive.ObjectID)
	AddComment(ctx context.Context, blogId primitive.ObjectID, request request.CreateCommentRequest, userId primitive.ObjectID) (response.BlogResponse, error)
	Like(ctx context.Context, blogId primitive.ObjectID, userId primitive.ObjectID) (response.BlogResponse, error)
	FindByTitle(ctx context.Context, title string, orderBy string, order string) []response.BlogResponse
}