package domain

import (
	"context"
	"errors"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
}

func NewUser(name, email string) *User {
	return &User{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

type UserRepository interface {
	Get(ctx context.Context, id int64) (user *User, err error)
}

var (
	ErrUserNotFound = errors.New("ユーザーが見つかりません。")
)
