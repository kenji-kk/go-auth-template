package data

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	UserName  string
	Email     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
