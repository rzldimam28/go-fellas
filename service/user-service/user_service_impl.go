package userservice

import (
	"context"

	"github.com/rzldimam28/wlb-test/model/entity"
	"github.com/rzldimam28/wlb-test/model/helper"
	"github.com/rzldimam28/wlb-test/model/web/request"
	"github.com/rzldimam28/wlb-test/model/web/response"
	userrepository "github.com/rzldimam28/wlb-test/repository/user-repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserServiceImpl struct {
	UserRepository userrepository.UserRepository
}

func NewUserService(userRepository userrepository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (userService *UserServiceImpl) FindAll(ctx context.Context) []response.UserResponse {
	users := userService.UserRepository.FindAll(ctx)
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponse := response.UserResponse{
			Id: user.Id,
			Username: user.Username,
			Email: user.Email,
			Status: user.Status,
		}
		userResponses = append(userResponses, userResponse)
	}
	return userResponses
}

func (userService *UserServiceImpl) FindById(ctx context.Context, userId primitive.ObjectID) response.UserResponse {
	user, err := userService.UserRepository.FindById(ctx, userId)
	helper.PanicIfError(err)

	userResponse := response.UserResponse{
		Id: user.Id,
		Username: user.Username,
		Email: user.Email,
		Status: user.Status,
	}
	return userResponse
}

func (userService *UserServiceImpl) FindByUsername(ctx context.Context, username string) response.UserResponse {
	user := userService.UserRepository.FindByUsername(ctx, username)
	userResponse := response.UserResponse{
		Id: user.Id,
		Username: user.Username,
		Email: user.Email,
		Status: user.Status,
	}
	return userResponse
}

func (userService *UserServiceImpl) Create(ctx context.Context, request request.CreateUserRequest) response.UserResponse {
	passwordHash, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	userToCreate := entity.Users{
		Username: request.Username,
		Email: request.Email,
		Password: passwordHash,
		Status: "Terdaftar",
	}
	user := userService.UserRepository.Create(ctx, userToCreate)
	userResponse := response.UserResponse{
		Id: user.Id,
		Username: user.Username,
		Email: user.Email,
		Status: user.Status,
	}
	return userResponse
}

func (userService *UserServiceImpl) Update(ctx context.Context, request request.UpdateUserRequest) response.UserResponse {
	userToUpdate, err := userService.UserRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	userToUpdate.Username = request.Username
	userToUpdate.Email = request.Email
	userToUpdate.Password = request.Password
	userToUpdate.Status = request.Status

	updatedUser := userService.UserRepository.Update(ctx, userToUpdate)
	userResponse := response.UserResponse{
		Id: updatedUser.Id,
		Username: updatedUser.Username,
		Email: updatedUser.Email,
		Status: updatedUser.Status,
	}
	return userResponse
}

func (userService *UserServiceImpl) Verify(ctx context.Context, userId primitive.ObjectID) response.UserResponse {
	userToUpdate, err := userService.UserRepository.FindById(ctx, userId)
	helper.PanicIfError(err)
	// userToUpdate.Status = request.Status

	updatedUser := userService.UserRepository.Verify(ctx, userToUpdate)
	userResponse := response.UserResponse{
		Id: updatedUser.Id,
		Username: updatedUser.Username,
		Email: updatedUser.Email,
		Status: "Aktif",
	}
	return userResponse
}

func (userService *UserServiceImpl) Delete(ctx context.Context, userId primitive.ObjectID) {
	userToDelete, err := userService.UserRepository.FindById(ctx, userId)
	helper.PanicIfError(err)
	userService.UserRepository.Delete(ctx, userToDelete)
}
