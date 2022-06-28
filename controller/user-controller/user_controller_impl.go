package usercontroller

import (
	"net/http"

	"github.com/rzldimam28/wlb-test/middleware"
	"github.com/rzldimam28/wlb-test/model/entity"
	"github.com/rzldimam28/wlb-test/model/helper"
	"github.com/rzldimam28/wlb-test/model/web/request"
	userservice "github.com/rzldimam28/wlb-test/service/user-service"
)

type UserControllerImpl struct {
	UserService userservice.UserService
}

func NewUserController(userService userservice.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (userController *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	users := userController.UserService.FindAll(r.Context())
	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Get All Users", users)
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	oid, err := helper.ReadParams(r, "userId")
	helper.PanicIfError(err)

	user := userController.UserService.FindById(r.Context(), oid)
	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Get User by Id", user)
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var userCreateReq request.CreateUserRequest
	helper.ReadFromRequestBody(r, &userCreateReq)

	userResp := userController.UserService.Create(r.Context(), userCreateReq)
	
	if userResp.Status == "Terdaftar" {
		err := helper.SendMailVerivication(userResp.Email, userResp.Id)
		if err != nil {
			webResponse := helper.CreateWebResponse(http.StatusNotFound, "Can not send email verifications", nil)
			helper.WriteToResponseBody(w, webResponse)
			return
		}
	}
	
	webResponse := helper.CreateWebResponse(http.StatusCreated, "Success Create New User. Please Verify your Email.", userResp)
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	oid, err := helper.ReadParams(r, "userId")
	helper.PanicIfError(err)
	
	var userUpdateReq request.UpdateUserRequest
	helper.ReadFromRequestBody(r, &userUpdateReq)

	userUpdateReq.Id = oid
	userResp := userController.UserService.Update(r.Context(), userUpdateReq)

	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Update User by Id", userResp)
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	oid, err := helper.ReadParams(r, "userId")
	helper.PanicIfError(err)
	
	userController.UserService.Delete(r.Context(), oid)

	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Delete Users", nil)
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var u entity.Credential
	helper.ReadFromRequestBody(r, &u)	

	userResponse := userController.UserService.FindByUsername(r.Context(), u.Username)
	if u.Username != userResponse.Username {
		webResponse := helper.CreateWebResponse(http.StatusUnauthorized, "Unauthorized", nil)
		helper.WriteToResponseBody(w, webResponse)
		return
	}
	
	if userResponse.Status != "Aktif" {
		webResponse := helper.CreateWebResponse(http.StatusUnauthorized, "Plis Verify You Email First", nil)
		helper.WriteToResponseBody(w, webResponse)
		return
	}

	token, err := middleware.GenerateToken(userResponse.Id)
	if err != nil {
		webResponse := helper.CreateWebResponse(http.StatusUnprocessableEntity, "Can not generate token", nil)
		helper.WriteToResponseBody(w, webResponse)
		return
	}
	webResponse := helper.CreateWebResponse(http.StatusOK, "Login Success. Please Use the Following Token.", token)
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) Verify(w http.ResponseWriter, r *http.Request) {
	oid, err := helper.ReadParams(r, "userId")
	helper.PanicIfError(err)

	userResp := userController.UserService.Verify(r.Context(), oid)

	webResponse := helper.CreateWebResponse(http.StatusOK, "Success Verify User", userResp)

	helper.WriteToResponseBody(w, webResponse)
}