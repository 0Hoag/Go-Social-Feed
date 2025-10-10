package comment

import "errors"

var wantErrors = []error{
	ErrRequiredField,
}

var (
	// follow
	ErrRequiredField   = errors.New("required field")
	ErrCommentNotFound = errors.New("Comment not found")
)
