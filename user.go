package flow

import (
	"context"
	"errors"
	"time"
)

var (
	errUserNameTaken = errors.New("username already taken")
	errEmailTaken    = errors.New("email already taken")
)

type UserRepository interface {
	Create(ctx context.Context, user User) (User, error)
	GetbyUsername(ctx context.Context, username string) (User, error)
	GetbyEmail(ctx context.Context, email string) (User, error)
}
type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
