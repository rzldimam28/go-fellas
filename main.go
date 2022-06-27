package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rzldimam28/wlb-test/config"
	blogcontroller "github.com/rzldimam28/wlb-test/controller/blog-contoller"
	usercontroller "github.com/rzldimam28/wlb-test/controller/user-controller"
	"github.com/rzldimam28/wlb-test/middleware"
	"github.com/rzldimam28/wlb-test/model/helper"
	blogrepository "github.com/rzldimam28/wlb-test/repository/blog-repository"
	commentrepository "github.com/rzldimam28/wlb-test/repository/comment-repository"
	userrepository "github.com/rzldimam28/wlb-test/repository/user-repository"
	blogservice "github.com/rzldimam28/wlb-test/service/blog-service"
	userservice "github.com/rzldimam28/wlb-test/service/user-service"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	fmt.Println("WLB Test")

	db := config.InitDatabase()

	// user
	userRepo := userrepository.NewUserRepository(db)
	userService := userservice.NewUserService(userRepo)
	userController := usercontroller.NewUserController(userService)

	// blog
	blogRepo := blogrepository.NewBlogRepository(db)
	commentRepo := commentrepository.NewCommentRepository(db)
	blogService := blogservice.NewBlogService(blogRepo, commentRepo, userRepo)
	blogController := blogcontroller.NewBlogController(blogService)

	r := mux.NewRouter()

	// blog
	blogRouter := r.PathPrefix("/blogs").Subrouter()
	blogRouter.HandleFunc("", blogController.FindAll).Methods("GET")
	blogRouter.HandleFunc("/{blogId}", blogController.FindById).Methods("GET")
	blogRouter.HandleFunc("", blogController.Create).Methods("POST")
	blogRouter.HandleFunc("/{blogId}", blogController.Update).Methods("PUT")
	blogRouter.HandleFunc("/{blogId}", blogController.Delete).Methods("DELETE")
	blogRouter.HandleFunc("/{blogId}/comments", blogController.AddComment).Methods("POST")
	blogRouter.HandleFunc("/{blogId}/like", blogController.Like).Methods("PUT")
	blogRouter.Use(middleware.Auth)

	// user
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/login", userController.Login).Methods("GET")
	userRouter.HandleFunc("/verify/{userId}", userController.Verify).Methods("PUT")
	userRouter.HandleFunc("", userController.FindAll).Methods("GET")
	userRouter.HandleFunc("/{userId}", userController.FindById).Methods("GET")
	userRouter.HandleFunc("", userController.Create).Methods("POST")
	userRouter.HandleFunc("/{userId}", userController.Update).Methods("PUT")
	userRouter.HandleFunc("/{userId}", userController.Delete).Methods("DELETE") 

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}