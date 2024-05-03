package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	UserName       string
	Email          string
	Password       *string
	Salt           []byte
	HashedPassword []byte
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

func NewUser(id uuid.UUID, name, email string, password *string, salt, hashedPassword []byte, createdAt, updatedAt *time.Time) (*User, error) {
	return &User{
		ID:             id,
		UserName:       name,
		Email:          email,
		Password:       password,
		Salt:           salt,
		HashedPassword: hashedPassword,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}, nil
}
