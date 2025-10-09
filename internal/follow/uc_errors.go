package follow

import "errors"

var wantErrors = []error{
	ErrRequiredField,
}

var (
	// follow
	ErrRequiredField  = errors.New("required field")
	ErrFollowNotFound = errors.New("follow not found")
)
