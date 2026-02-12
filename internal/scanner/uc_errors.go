package scanner

import "errors"

var wantErrors = []error{
	ErrTokenNotFound,
	ErrSourceCodeNotFound,
}

var (
	// Scanner errors
	ErrTokenNotFound      = errors.New("token not found on DexScreener")
	ErrSourceCodeNotFound = errors.New("source code not found on supported networks")
)
