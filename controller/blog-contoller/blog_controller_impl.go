package blogcontoller

import (
	"net/http"

	"github.com/rzldimam28/wlb-test/model/helper"
	"github.com/rzldimam28/wlb-test/model/web/request"
	blogservice "github.com/rzldimam28/wlb-test/service/blog-service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogControllerImpl struct {
	BlogService blogservice.BlogService
}

func NewBlogController(blogService blogservice.BlogService) BlogController {
	return &BlogControllerImpl{
		BlogService: blogService,
	}
}

func (blogController *BlogControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Query().Get("title")
	orderBy := r.URL.Query().Get("order_by")
	asc := r.URL.Query().Get("ascending")

	if title != "" {
		blogs := blogController.BlogService.FindByTitle(r.Context(), title, orderBy, asc)
		webResponse := helper.CreateWebResponse(http.StatusOK, "Success Get All Blogs", blogs)
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	blogs := blogController.BlogService.FindAll(r.Context())
	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Get All Blogs", blogs)
	helper.WriteToResponseBody(w, webResponse)
}

func (blogController *BlogControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	oid, err := helper.ReadParams(r, "blogId")
	helper.PanicIfError(err)

	blog := blogController.BlogService.FindById(r.Context(), oid)
	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Get Blog by Id", blog)
	helper.WriteToResponseBody(w, webResponse)
}

func (blogController *BlogControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(primitive.ObjectID)
	
	var blogCreateReq request.CreateBlogRequest
	helper.ReadFromRequestBody(r, &blogCreateReq)
	
	blogResp := blogController.BlogService.Create(r.Context(), blogCreateReq, userId)
	webResponse := helper.CreateWebResponse(http.StatusCreated, "Success Create New Blog", blogResp)

	helper.WriteToResponseBody(w, webResponse)
}

func (blogController *BlogControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	oid, err := helper.ReadParams(r, "blogId")
	helper.PanicIfError(err)
	
	var blogUpdateReq request.UpdateBlogRequest
	helper.ReadFromRequestBody(r, &blogUpdateReq)

	blogUpdateReq.Id = oid
	blogResp := blogController.BlogService.Update(r.Context(), blogUpdateReq)

	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Update Blog by Id", blogResp)
	helper.WriteToResponseBody(w, webResponse)
}

func (blogController *BlogControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	oid, err := helper.ReadParams(r, "blogId")
	helper.PanicIfError(err)
	
	blogController.BlogService.Delete(r.Context(), oid)

	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Delete Blog", nil)
	helper.WriteToResponseBody(w, webResponse)
}

func (blogController *BlogControllerImpl) AddComment(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(primitive.ObjectID)

	blogId, err := helper.ReadParams(r, "blogId")
	helper.PanicIfError(err)

	var blogAddComment request.CreateCommentRequest
	helper.ReadFromRequestBody(r, &blogAddComment)

	blogResp, err := blogController.BlogService.AddComment(r.Context(), blogId, blogAddComment, userId)
	if err != nil {
		webResponse := helper.CreateWebResponse(http.StatusUnauthorized, err.Error(), nil)
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Add Comments", blogResp)
	helper.WriteToResponseBody(w, webResponse)
}

func (blogController *BlogControllerImpl) Like(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(primitive.ObjectID)
	
	oid, err := helper.ReadParams(r, "blogId")
	helper.PanicIfError(err)
	
	blogResp, err := blogController.BlogService.Like(r.Context(), oid, userId)
	if err != nil {
		webResponse := helper.CreateWebResponse(http.StatusBadRequest, err.Error(), nil)
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Update Blog by Id", blogResp)
	helper.WriteToResponseBody(w, webResponse)
}