package graph

import (
	"context"
	"errors"

	"github.com/arisromil/flow"
)

func mapAuthResponse(res flow.AuthResponse) *AuthResponse {
	return &AuthResponse{
		AccessToken: res.AccessToken,
		User:        mapUser(res.User),
	}
}

func (r *mutationResolver) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {
	res, err := r.AuthService.Register(ctx, flow.RegisterInput{
		Email:           input.Email,
		Username:        input.Username,
		Password:        input.Password,
		ConfirmPassword: input.Password,
	})

	if err != nil {
		switch {
		case errors.Is(err, flow.NewValidationError) ||
			errors.Is(err, flow.ErrBadCredentials):
			return nil, buildBadRequestError(ctx, err)
		default:
			return nil, err
		}
	}

	return mapAuthResponse(res), nil
}

func (r *mutationResolver) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	panic("not implemented")
}
