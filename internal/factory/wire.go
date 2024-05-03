//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/kenji-kk/go-auth-template/internal/applicationService"
	"github.com/kenji-kk/go-auth-template/internal/handler"
	"github.com/kenji-kk/go-auth-template/internal/infra"
	"github.com/kenji-kk/go-auth-template/internal/repository"

	"github.com/google/wire"
)

func InitializeRootHandlers() handler.RootHandlers {
	wire.Build(
		// handler
		handler.NewRootHandlers,
		handler.NewUserHandler,

		// usecase
		applicationService.NewUserApplicationService,

		// repository
		repository.NewUserRepository,

		// DB
		infra.NewMysql,
	)
	return handler.RootHandlers{}
}
