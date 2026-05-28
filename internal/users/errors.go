package users

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrEmptyEmail   = errors.New("email cannot be empty")
)
