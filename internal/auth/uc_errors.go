package auth

import "errors"

var wantErrors = []error{
	ErrUserNotFound,
	ErrRequiredField,
}

var (
	// auth
	ErrUserNotFound  = errors.New("user not found")
	ErrInvalidCreds  = errors.New("invalid phone or password")
	ErrRequiredField = errors.New("required field")
)
