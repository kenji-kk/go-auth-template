package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/kenji-kk/go-auth-template/internal/factory"
	"github.com/kenji-kk/go-auth-template/internal/handler"
	"github.com/kenji-kk/go-auth-template/pkg/logger"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	return &Server{echo: echo.New()}
}

func (s *Server) Run() error {
	// rootHandlers生成
	rootHandlers := factory.InitializeRootHandlers()

	s.MapHandler(&rootHandlers)

	logger.Logger.Info("server start")
	s.echo.Logger.Fatal(s.echo.Start(":8080"))

	return nil
}

func (s *Server) MapHandler(rootHandlers *handler.RootHandlers) error {
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.Logger())

	s.echo.POST("/users", rootHandlers.UserHandler.Create)
	s.echo.POST("/login", rootHandlers.UserHandler.Login)

	return nil
}
