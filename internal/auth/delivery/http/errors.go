package http

import (
	"github.com/hoag/go-social-feed/internal/users"
	pkgErrors "github.com/hoag/go-social-feed/pkg/errors"
)

var (
	errWrongBody = pkgErrors.NewHTTPError(140003, "Wrong body")

	// User errors
	errUserNotFound = pkgErrors.NewHTTPError(143005, "User not found")
)

func (h handler) mapError(err error) error {
	switch err {
	case users.ErrUserNotFound:
		return errUserNotFound
	default:
		panic(err)
	}
}
