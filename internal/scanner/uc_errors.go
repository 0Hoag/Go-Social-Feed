package scanner

import "errors"

var wantErrors = []error{
	ErrScanToken,
}

var (
	// user
	ErrScanToken = errors.New("user not found")
)
