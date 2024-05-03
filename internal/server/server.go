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

	// 認証が必要なルート
	// s.echo.GET("/secure", secureEndpoint, echojwt.WithConfig(echojwt.Config{
	// 	SigningKey:  []byte("your_secret_key"),
	// 	TokenLookup: "cookie:jwt", // クッキーからJWTを取得
	// }))
	// エンドポイントの例
	// func secureEndpoint(c echo.Context) error {
	// 	user := c.Get("user").(*jwt.Token)
	// 	claims := user.Claims.(jwt.MapClaims)
	// 	userID := claims["user_id"].(string) // ユーザーIDの取得
	// 	return c.String(http.StatusOK, "Access granted for user ID: "+userID)
	// }

	return nil
}
