package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/arisromil/flow"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository flow.UserRepository
}

func NewAuthService(userRepository flow.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) Register(ctx context.Context, input flow.RegisterInput) (flow.AuthResponse, error) {
	input.Sanitize()
	if err := input.Validate(); err != nil {
		return flow.AuthResponse{}, err
	}

	if _, err := s.UserRepository.GetbyUsername(ctx, input.Username); !errors.Is(err, flow.ErrNotFound) {
		return flow.AuthResponse{}, flow.ErrUserNameTaken
	}

	var (
		errUserNameTaken = errors.New("username is already taken")
	)

	if _, err := s.UserRepository.GetbyEmail(ctx, input.Email); !errors.Is(err, flow.ErrNotFound) {
		return flow.AuthResponse{}, flow.ErrEmailTaken
	}

	user := flow.User{
		Username: input.Username,
		Email:    input.Email,
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return flow.AuthResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = string(hashPassword)

	user, err = s.UserRepository.Create(ctx, user)
	if err != nil {
		return flow.AuthResponse{}, fmt.Errorf("failed to create user: %w", err)
	}

	return flow.AuthResponse{
		AccessToken: "dummy_token",
		User:        user,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, input flow.LoginInput) (flow.AuthResponse, error) {
	input.Sanitize()
	if err := input.Validate(); err != nil {
		return flow.AuthResponse{}, err
	}
	user, err := s.UserRepository.GetbyEmail(ctx, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, flow.ErrNotFound):
			return flow.AuthResponse{}, flow.ErrBadCredentiaLS
		default:
			return flow.AuthResponse{}, err
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return flow.AuthResponse{}, flow.ErrBadCredentiaLS
	}
	return flow.AuthResponse{
		AccessToken: "dummy_token",
		User:        user,
	}, nil
}
