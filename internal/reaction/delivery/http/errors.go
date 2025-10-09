package http

import (
	"github.com/hoag/go-social-feed/internal/reaction"
	pkgErrors "github.com/hoag/go-social-feed/pkg/errors"
)

var (
	errWrongPaginationQuery = pkgErrors.NewHTTPError(140001, "Wrong pagination query")
	errWrongQuery           = pkgErrors.NewHTTPError(140002, "Wrong query")
	errWrongBody            = pkgErrors.NewHTTPError(140003, "Wrong body")

	// Reaction errors
	errReactionNotFound = pkgErrors.NewHTTPError(143004, "Reaction not found")
)

func (h handler) mapError(err error) error {
	switch err {
	case reaction.ErrReactionNotFound:
		return errReactionNotFound
	default:
		panic(err)
	}
}
