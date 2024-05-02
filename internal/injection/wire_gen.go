// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/kenji-kk/go-auth-template/internal/handler"
	"github.com/kenji-kk/go-auth-template/internal/infra"
	"github.com/kenji-kk/go-auth-template/internal/repository"
	"github.com/kenji-kk/go-auth-template/internal/usecase"
)

// Injectors from wire.go:

func InitializeRootHandlers() handler.RootHandlers {
	db := infra.NewMysql()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	rootHandlers := handler.NewRootHandlers(userHandler)
	return rootHandlers
}