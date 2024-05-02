//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/kenji-kk/go-auth-template/internal/handler"
	"github.com/kenji-kk/go-auth-template/internal/infra"
	"github.com/kenji-kk/go-auth-template/internal/repository"
	"github.com/kenji-kk/go-auth-template/internal/usecase"

	"github.com/google/wire"
)

func InitializeRootHandlers() handler.RootHandlers {
	wire.Build(
		// handler
		handler.NewRootHandlers,
		handler.NewUserHandler,

		// usecase
		usecase.NewUserUsecase,

		// repository
		repository.NewUserRepository,

		// DB
		infra.NewMysql,
	)
	return handler.RootHandlers{}
}
