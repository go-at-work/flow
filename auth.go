package flow

import (
	"context"
	"fmt"
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

type AuthService interface {
	Register(ctx context.Context, input RegisterInput) (AuthResponse error)
}

type AuthResponse struct {
	AccessToken string
	User        User
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
