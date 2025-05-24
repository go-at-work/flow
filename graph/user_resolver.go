package graph

import (
	"context"

	"github.com/arisromil/flow"
)

func mapUser(user flow.User) *User {
	return &User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func (r *queryResolver) Me(ctx context.Context) (*User, error) {
	panic("not implemented")
}
