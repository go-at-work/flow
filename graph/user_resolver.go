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
	userId, err := flow.GetUserIdFromContext(ctx)
	if err != nil {
		return nil, flow.ErrUnAuthenticated
	}
	return mapUser(flow.User{
		ID: userId,
	}), nil

}
