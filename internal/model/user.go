package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/kenji-kk/go-auth-template/internal/domain"
)

type User struct {
	ID             uuid.UUID
	UserName       string
	Email          string
	Salt           []byte
	HashedPassword []byte
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

func NewUser(domainUser *domain.User) (*User, error) {
	return &User{
		ID:             domainUser.ID,
		UserName:       domainUser.UserName,
		Email:          domainUser.Email,
		Salt:           domainUser.Salt,
		HashedPassword: domainUser.HashedPassword,
		CreatedAt:      domainUser.CreatedAt,
		UpdatedAt:      domainUser.UpdatedAt,
	}, nil
}
