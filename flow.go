package flow

import "errors"

var (
	ErrBadCredentiaLS  = errors.New("bad credentials")
	ErrNotFound        = errors.New("user not found")
	NewValidationError = errors.New("validation error")
)
