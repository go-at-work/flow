package flow

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

var (
	UsenameMinLength  = 2
	PasswordMinLength = 6
)

type RegisterInput struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

type LoginInput struct {
	Email    string
	Password string
}

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (AuthResponse, error)
	Login(ctx context.Context, input LoginInput) (AuthResponse, error)
}

type AuthTokenService interface {
	CreateToken(ctx context.Context, user User) (string, error)
	CreateRefreshToken(ctx context.Context, user User, tokenId string) (string, error)
	ParseToken(ctx context.Context, payload string) (AuthToken, error)
	ParseTokenFromRequest(ctx context.Context, r *http.Request) (AuthToken, error)
}

type AuthToken struct {
	ID  string
	Sub string
}

type AuthResponse struct {
	AccessToken string
	User        User
}

func (in LoginInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)
}

func (in LoginInput) Validate() error {

	if len(in.Password) < 1 {
		return fmt.Errorf("%w: password required", NewValidationError)
	}

	return nil

}

func (in RegisterInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)
	in.Username = strings.TrimSpace(in.Username)
}

func (in RegisterInput) Validate() error {
	if len(in.Username) < UsenameMinLength {
		return fmt.Errorf("%w: username must be at least (%d) characters long", NewValidationError, UsenameMinLength)

	}
	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf("%w: password must be at least (%d) characters long", NewValidationError, PasswordMinLength)
	}
	if in.Password != in.ConfirmPassword {
		return fmt.Errorf("%w: confirm password must match the password", NewValidationError)
	}
	return nil

}
