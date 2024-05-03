package applicationService

import (
	"context"
	"crypto/rand"

	"github.com/google/uuid"
	"github.com/kenji-kk/go-auth-template/internal/domain"
	"github.com/kenji-kk/go-auth-template/internal/dto/data"
	"github.com/kenji-kk/go-auth-template/internal/dto/request"
	"github.com/kenji-kk/go-auth-template/internal/model"
	"github.com/kenji-kk/go-auth-template/internal/repository"
	"github.com/kenji-kk/go-auth-template/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserApplicationService interface {
	Create(context.Context, *request.CreateUser) (*data.User, error)
}

type userApplicationService struct {
	userRepository repository.UserRepository
}

func NewUserApplicationService(userRepository repository.UserRepository) UserApplicationService {
	return &userApplicationService{userRepository}
}

func (userApplicationService *userApplicationService) Create(ctx context.Context, createInput *request.CreateUser) (*data.User, error) {
	uuid := uuid.New()

	salt, err := GenerateSalt()
	if err != nil {
		return nil, err
	}

	// hashPassword作成
	toHash := append([]byte(createInput.Password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(toHash, bcrypt.DefaultCost)
	if err != nil {
		logger.Logger.Error("An error occurred while creating the hashedPassword", zap.Error(err))
		return nil, err
	}

	userDomain, err := domain.NewUser(
		uuid,
		createInput.UserName,
		createInput.Email,
		&createInput.Password,
		salt,
		hashedPassword,
		nil,
		nil,
	)
	if err != nil {
		logger.Logger.Error("An error occurred while creating userDomain", zap.Error(err))
		return nil, err
	}

	userModel, err := model.NewUser(userDomain)
	if err != nil {
		logger.Logger.Error("An error occurred while creating userModel", zap.Error(err))
		return nil, err
	}

	createdUser, err := userApplicationService.userRepository.Create(ctx, userModel)
	if err != nil {
		logger.Logger.Error("An error occurred while userUsecase.userRepository.CreateUser", zap.Error(err))
		return nil, err
	}

	return &data.User{
		ID:        createdUser.ID,
		UserName:  createdUser.UserName,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, err
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		logger.Logger.Error("An error occurred while creating the salt", zap.Error(err))
		return nil, err
	}
	return salt, nil
}
