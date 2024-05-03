package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kenji-kk/go-auth-template/internal/applicationService"
	"github.com/kenji-kk/go-auth-template/internal/const/http/cookie"
	"github.com/kenji-kk/go-auth-template/internal/dto/request"
	"github.com/kenji-kk/go-auth-template/pkg/utils"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Create(echo.Context) error
	Login(c echo.Context) error
}

type userHandler struct {
	userApplicationService applicationService.UserApplicationService
}

func NewUserHandler(userApplicationService applicationService.UserApplicationService) UserHandler {
	return &userHandler{userApplicationService}
}

func (userHandler *userHandler) Create(c echo.Context) error {
	ctx := context.Background()
	requestCreateUser := new(request.CreateUser)
	if err := utils.ReadRequest(c, requestCreateUser); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createdUser, err := userHandler.userApplicationService.Create(ctx, requestCreateUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	jws := createJWT(createdUser.ID.String())

	utils.WriteCookie(c, cookie.JWSCookieName, jws, 1)

	return c.JSON(http.StatusCreated, createdUser)
}

func createJWT(userID string) string {
	// Claimsオブジェクト作成
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダーとペイロード作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	jws, _ := token.SignedString([]byte("your_secret_key"))

	return jws
}

func (userHandler *userHandler) Login(c echo.Context) error {
	ctx := context.Background()
	requestLogin := new(request.Login)
	if err := utils.ReadRequest(c, requestLogin); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	loginUser, err := userHandler.userApplicationService.Login(ctx, requestLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	jws := createJWT(loginUser.ID.String())

	utils.WriteCookie(c, cookie.JWSCookieName, jws, 1)

	return c.JSON(http.StatusOK, loginUser)
}
