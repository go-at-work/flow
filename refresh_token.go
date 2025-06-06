package flow

import (
	"context"
	"time"
)

var (
	AccessTokenLifetime  = time.Minute * 15   // 15 minutes
	RefreshTokenLifetime = time.Hour * 24 * 7 // 1 week
)

type RefreshToken struct {
	ID         string
	Name       string
	UserID     string
	LastUsedAt time.Time
	ExpiresAt  time.Time
	CreatedAt  time.Time
}

type RefreshTokenParams struct {
	Sub  string
	Name string
}

type RefreshTokenRepo interface {
	Create(ctx context.Context, params RefreshTokenParams) (RefreshToken, error)
	GetByID(ctx context.Context, id string) (RefreshToken, error)
}
