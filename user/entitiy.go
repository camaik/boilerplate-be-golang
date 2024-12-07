package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:binary(16);primary_key"`
	Name           string
	Email          string `gorm:"unique"`
	PasswordHash   string
	AvatarFileName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
