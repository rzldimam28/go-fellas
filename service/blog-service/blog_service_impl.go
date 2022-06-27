package blogservice

import (
	"context"
	"errors"

	"github.com/rzldimam28/wlb-test/model/entity"
	"github.com/rzldimam28/wlb-test/model/helper"
	"github.com/rzldimam28/wlb-test/model/web/request"
	"github.com/rzldimam28/wlb-test/model/web/response"
	blogrepository "github.com/rzldimam28/wlb-test/repository/blog-repository"
	commentrepository "github.com/rzldimam28/wlb-test/repository/comment-repository"
	userrepository "github.com/rzldimam28/wlb-test/repository/user-repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogServiceImpl struct {
	BlogRepository blogrepository.BlogRepository
	UserRepository userrepository.UserRepository
	CommentRepository commentrepository.CommentRepository
}

func NewBlogService(blogRepository blogrepository.BlogRepository, commentRepository commentrepository.CommentRepository, userRepository userrepository.UserRepository) BlogService {
	return &BlogServiceImpl{
		BlogRepository: blogRepository,
		CommentRepository: commentRepository,
		UserRepository: userRepository,
	}
}

func (blogService *BlogServiceImpl) FindAll(ctx context.Context) []response.BlogResponse {
	blogs := blogService.BlogRepository.FindAll(ctx)
	var blogResponses []response.BlogResponse
	for _, blog := range blogs {
		com := blogService.CommentRepository.FindByBlogId(ctx, blog.Id)
		blogResponse := response.BlogResponse{
			Id: blog.Id,
			UserId: blog.UserId,
			Title: blog.Title,
			Content: blog.Content,
			Comments: &com,
			IsCom: blog.IsCom,
			LikedBy: blog.LikedBy,
			LikedCount: blog.LikedCount,
		}
		blogResponses = append(blogResponses, blogResponse)
	}
	return blogResponses
}

func (blogService *BlogServiceImpl) FindById(ctx context.Context, blogId primitive.ObjectID) response.BlogResponse {
	blog, err := blogService.BlogRepository.FindById(ctx, blogId)
	helper.PanicIfError(err)

	com := blogService.CommentRepository.FindByBlogId(ctx, blog.Id)
	blogResponse := response.BlogResponse{
		Id: blog.Id,
		UserId: blog.UserId,
		Title: blog.Title,
		Content: blog.Content,
		Comments: &com,
		IsCom: blog.IsCom,
		LikedBy: blog.LikedBy,
		LikedCount: blog.LikedCount,
	}
	return blogResponse
}

func (blogService *BlogServiceImpl) Create(ctx context.Context, request request.CreateBlogRequest, id primitive.ObjectID) response.BlogResponse {
	blogToCreate := entity.Blogs{
		UserId: id,
		Title: request.Title,
		Content: request.Content,
		IsCom: request.IsCom,
	}
	blog := blogService.BlogRepository.Create(ctx, blogToCreate)
	blogResponse := response.BlogResponse{
		Id: blog.Id,
		UserId: blog.UserId,
		Title: blog.Title,
		Content: blog.Content,
		Comments: blog.Comment,
		IsCom: blog.IsCom,
		LikedBy: blog.LikedBy,
		LikedCount: blog.LikedCount,
	}
	return blogResponse
}

func (blogService *BlogServiceImpl) Update(ctx context.Context, request request.UpdateBlogRequest) response.BlogResponse {
	blogToUpdate, err := blogService.BlogRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	blogToUpdate.Title = request.Title
	blogToUpdate.Content = request.Content

	updatedBlog := blogService.BlogRepository.Update(ctx, blogToUpdate)
	blogResponse := response.BlogResponse{
		Id: updatedBlog.Id,
		UserId: updatedBlog.UserId,
		Title: updatedBlog.Title,
		Content: updatedBlog.Content,
		Comments: updatedBlog.Comment,
		IsCom: updatedBlog.IsCom,
		LikedBy: updatedBlog.LikedBy,
		LikedCount: updatedBlog.LikedCount,
	}
	return blogResponse
}

func (blogService *BlogServiceImpl) Delete(ctx context.Context, blogId primitive.ObjectID) {
	blogToDelete, err := blogService.BlogRepository.FindById(ctx, blogId)
	helper.PanicIfError(err)

	blogService.BlogRepository.Delete(ctx, blogToDelete)
}

func (blogService *BlogServiceImpl) AddComment(ctx context.Context, blogId primitive.ObjectID, request request.CreateCommentRequest, userId primitive.ObjectID) (response.BlogResponse, error) {	
	commentToAdd := entity.Comments{
		UserId: userId,
		BlogId: blogId,
		Content: request.Content,
	}	

	blogs, err := blogService.BlogRepository.FindById(ctx, blogId)
	helper.PanicIfError(err)
	
	if !blogs.IsCom {
		var blogResp response.BlogResponse
		return blogResp, errors.New("this blog is uncommentable")
	}
	
	userBlogId, err := blogService.UserRepository.FindById(ctx, blogs.UserId)
	helper.PanicIfError(err)

	err = helper.SendMailCommentNotif(userBlogId.Email)
	if err != nil {
		var blogResp response.BlogResponse
		return blogResp, err
	}

	commentToAdd = blogService.CommentRepository.Insert(ctx, commentToAdd)	
	allComments := blogService.CommentRepository.FindByBlogId(ctx, blogId)
	
	blogs.Comment = &allComments

	blogResp := response.BlogResponse{
		Id: blogs.Id,
		UserId: blogs.UserId,
		Title: blogs.Title,
		Content: blogs.Content,
		Comments: blogs.Comment,
		IsCom: blogs.IsCom,
		LikedBy: blogs.LikedBy,
		LikedCount: blogs.LikedCount,
	}
	return blogResp, nil
}

func (blogService *BlogServiceImpl) Like(ctx context.Context, blogId primitive.ObjectID, userId primitive.ObjectID) (response.BlogResponse, error) {
	blogToUpdate, err := blogService.BlogRepository.FindById(ctx, blogId)
	helper.PanicIfError(err)

	for _, user := range blogToUpdate.LikedBy {
		if user == userId {
			var blogResponse response.BlogResponse
			return blogResponse, errors.New("this users already like this blog")
		}
	}

	blogToUpdate.LikedCount += 1
	blogToUpdate.LikedBy = append(blogToUpdate.LikedBy, userId)

	userBlogId, err := blogService.UserRepository.FindById(ctx, blogToUpdate.UserId)
	helper.PanicIfError(err)

	err = helper.SendMailLikeNotif(userBlogId.Email)
	if err != nil {
		var blogResp response.BlogResponse
		return blogResp, err
	}

	updatedBlog := blogService.BlogRepository.Like(ctx, blogToUpdate)
	blogResponse := response.BlogResponse{
		Id: updatedBlog.Id,
		UserId: updatedBlog.UserId,
		Title: updatedBlog.Title,
		Content: updatedBlog.Content,
		Comments: updatedBlog.Comment,
		IsCom: updatedBlog.IsCom,
		LikedBy: updatedBlog.LikedBy,
		LikedCount: updatedBlog.LikedCount,
	}

	return blogResponse, nil
}

func (blogService *BlogServiceImpl) FindByTitle(ctx context.Context, title string, orderBy string, order string) []response.BlogResponse {
	
	var orderInt int
	if order == "" || order == "false" {
		orderInt = 1
	} else {
		orderInt = -1
	}

	blogs, err := blogService.BlogRepository.FindByTitle(ctx, title, orderBy, orderInt)
	helper.PanicIfError(err)

	var blogResponses []response.BlogResponse
	for _, blog := range blogs {
		com := blogService.CommentRepository.FindByBlogId(ctx, blog.Id)
		blogResponse := response.BlogResponse{
			Id: blog.Id,
			UserId: blog.UserId,
			Title: blog.Title,
			Content: blog.Content,
			Comments: &com,
			IsCom: blog.IsCom,
			LikedBy: blog.LikedBy,
			LikedCount: blog.LikedCount,
		}
		blogResponses = append(blogResponses, blogResponse)
	}
	return blogResponses
}