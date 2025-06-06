package flow

import "errors"

var (
	ErrBadCredentials     = errors.New("bad credentials")
	ErrNotFound           = errors.New("user not found")
	NewValidationError    = errors.New("validation error")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrNoUserIdInContext  = errors.New("no user id in context")
	ErrGenAccessToken     = errors.New("failed to generate access token")
	ErrUnAuthenticated    = errors.New("unauthenticated")
)
