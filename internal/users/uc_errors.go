package users

import "errors"

var wantErrors = []error{
	ErrUserNotFound,
	ErrRequiredField,
}

var (
	// user
	ErrUserNotFound  = errors.New("user not found")
	ErrRequiredField = errors.New("required field")
)
