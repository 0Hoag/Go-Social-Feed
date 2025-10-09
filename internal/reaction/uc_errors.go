package reaction

import "errors"

var wantErrors = []error{
	ErrRequiredField,
}

var (
	// reaction
	ErrRequiredField    = errors.New("required field")
	ErrReactionNotFound = errors.New("reaction not found")
)
