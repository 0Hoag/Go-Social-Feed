package http

import (
	"github.com/hoag/go-social-feed/internal/follow"
	pkgErrors "github.com/hoag/go-social-feed/pkg/errors"
)

var (
	errWrongPaginationQuery = pkgErrors.NewHTTPError(140001, "Wrong pagination query")
	errWrongQuery           = pkgErrors.NewHTTPError(140002, "Wrong query")
	errWrongBody            = pkgErrors.NewHTTPError(140003, "Wrong body")

	// Follow errors
	errFollowNotFound = pkgErrors.NewHTTPError(143004, "Follow not found")
)

func (h handler) mapError(err error) error {
	switch err {
	case follow.ErrFollowNotFound:
		return errFollowNotFound
	default:
		panic(err)
	}
}
