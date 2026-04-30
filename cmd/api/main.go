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
	"techconnect/internal/domain/user"
)

func main() {
	cfg := config.Load()

	hasher := security.BcryptHasher{}
	userRepo := postgres.NewUserRepository()
	userSvc := user.NewService(userRepo, hasher)
	registerUC := usecase.NewRegisterUseCase(userSvc)

	userHandler := httpAdapter.NewUserHandler(registerUC)
	postHandler := httpAdapter.NewPostHandler()

	router := httpAdapter.NewRouter(userHandler, postHandler)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("servidor rodando em %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
