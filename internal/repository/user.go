package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/kenji-kk/go-auth-template/internal/model"
)

type UserRepository interface {
	Create(context.Context, *model.User) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if res := userRepository.db.Create(&user); res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
