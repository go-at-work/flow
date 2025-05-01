package domain

import (
	"context"
	"testing"

	"github.com/arisromil/flow"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAuthService_Register(t *testing.T) {
	validInput := flow.RegisterInput{
		Username:        "validuser",
		Email:           "validuser@example.com",
		Password:        "validpassword",
		ConfirmPassword: "validpassword",
	}

	t.Run("can register", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.userRepo{}

		userRepo.On("GetbyUsername", mock.Anything, mock.Anything).Return(flow.User{}, error.ErrNotFound)

		userRepo.On("GetbyEmail", mock.Anything, mock.Anything).Return(flow.User{}, error.ErrNotFound)

		userRepo.On("Create", mock.Anything, mock.Anything).Return(flow.User{ID: "123",
			Username: validInput.Username,
			Email:    validInput.Email}, nil)

		service := NewAuthService(userRepo)

		res, err := service.Register(context.Background(), validInput)

		require.NoError(t, err)

		require.NotEmpty(t, res.AccessToken)
		require.NotEmpty(t, res.User.ID)
		require.NotEmpty(t, res.User.Username)
		require.NotEmpty(t, res.User.Email)

		userRepo.AssertExpectations(t)

	})
}
