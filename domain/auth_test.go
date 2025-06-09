package domain

import (
	"context"
	"errors"
	"testing"

	"github.com/arisromil/flow"
	"github.com/arisromil/flow/faker"
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

		authTokenService := &mocks.authTokenService{}

		authTokenService.On("CreateToken", mock.Anything, mock.Anything).Return("validaccesstoken", nil)

		service := NewAuthService(userRepo, authTokenService)

		res, err := service.Register(context.Background(), validInput)

		require.NoError(t, err)

		require.NotEmpty(t, res.AccessToken)
		require.NotEmpty(t, res.User.ID)
		require.NotEmpty(t, res.User.Username)
		require.NotEmpty(t, res.User.Email)

		userRepo.AssertExpectations(t)

	})
}

func TestAuthService_Login(t *testing.T) {

	validInput := flow.LoginInput{
		Username: "validuser",
		Password: "validpassword",
	}

	t.Run("can login", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.userRepo{}

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).Return(flow.User{
			Email:    validInput.Email,
			Password: faker.Password,
		}, nil)

		authTokenService := &mocks.authTokenService{}
		authTokenService.On("CreateToken", mock.Anything, mock.Anything).Return("validaccesstoken", nil)
		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Login(ctx, validInput)
		require.NoError(t, err)

		userRepo.AssertExpectations(t)

	})

	t.Run("wrong password", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.userRepo{}

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).Return(flow.User{
			Email:    validInput.Email,
			Password: faker.Password,
		}, nil)

		authTokenService := &mocks.authTokenService{}
		service := NewAuthService(userRepo, authTokenService)

		validInput.Password = "wrongpasswordsomething"

		_, err := service.Login(ctx, validInput)
		require.ErrorIs(t, err, flow.ErrBadCredentials)

		userRepo.AssertExpectations(t)

	})

	t.Run("email not found", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.userRepo{}

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).Return(flow.User{}, flow.ErrNotFound)

		authTokenService := &mocks.authTokenService{}
		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Login(ctx, validInput)
		require.ErrorIs(t, err, flow.ErrBadCredentials)

		userRepo.AssertExpectations(t)

	})

	t.Run("get user by email", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.userRepo{}

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).Return(flow.User{}, errors.New("something went wrong"))

		authTokenService := &mocks.authTokenService{}
		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Login(ctx, validInput)
		require.Error(t, err)

		userRepo.AssertExpectations(t)

	})

	t.Run("invalid input", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.userRepo{}

		authTokenService := &mocks.authTokenService{}
		service := NewAuthService(userRepo, authTokenService)

		_, err := service.Login(ctx, flow.LoginInput{
			Email:    "invalidemail",
			Password: "",
		})
		require.Error(t, err, flow.ErrBadCredentials)

		userRepo.AssertExpectations(t)

	})

}
