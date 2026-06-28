package main

import (
	"fmt"
	"log"
	"net/http"
	"techconnect/config"
	httpAdapter "techconnect/internal/adapter/in/http"
	"techconnect/internal/adapter/out/postgres"
	"techconnect/internal/adapter/out/security"
	"techconnect/internal/application/usecase"
	"techconnect/internal/domain/post"
	"techconnect/internal/domain/user"
)

func main() {
	cfg := config.Load()
	db := config.NewDB(cfg.DatabaseURL)

	hasher := security.BcryptHasher{}

	// User
	userRepo := postgres.NewUserRepository(db)
	userSvc := user.NewService(userRepo, hasher)
	registerUC := usecase.NewRegisterUseCase(userSvc)
	getUserUC := usecase.NewGetUserUseCase(userSvc)
	updateUserUC := usecase.NewUpdateUserUseCase(userSvc)
	deleteUserUC := usecase.NewDeleteUserUseCase(userSvc)

	// Post + Comment
	postRepo := postgres.NewPostRepository(db)
	commentRepo := postgres.NewCommentRepository(db)
	postSvc := post.NewService(postRepo, commentRepo)
	createPostUC := usecase.NewCreatePostUseCase(postSvc)
	getPostUC := usecase.NewGetPostUseCase(postSvc)
	listPostsUC := usecase.NewListPostsUseCase(postSvc)
	deletePostUC := usecase.NewDeletePostUseCase(postSvc)
	addCommentUC := usecase.NewAddCommentUseCase(postSvc)
	deleteCommentUC := usecase.NewDeleteCommentUseCase(postSvc)

	userHandler := httpAdapter.NewUserHandler(registerUC, getUserUC, updateUserUC, deleteUserUC)
	postHandler := httpAdapter.NewPostHandler(createPostUC, getPostUC, listPostsUC, deletePostUC, addCommentUC, deleteCommentUC)

	router := httpAdapter.NewRouter(userHandler, postHandler)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("servidor rodando em %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
