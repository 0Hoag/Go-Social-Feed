package http

import (
	"github.com/hoag/go-social-feed/internal/comment"
	pkgErrors "github.com/hoag/go-social-feed/pkg/errors"
)

var (
	errWrongPaginationQuery = pkgErrors.NewHTTPError(140001, "Wrong pagination query")
	errWrongQuery           = pkgErrors.NewHTTPError(140002, "Wrong query")
	errWrongBody            = pkgErrors.NewHTTPError(140003, "Wrong body")

	// Comment errors
	errCommentNotFound = pkgErrors.NewHTTPError(143004, "Comment not found")
)

func (h handler) mapError(err error) error {
	switch err {
	case comment.ErrCommentNotFound:
		return errCommentNotFound
	default:
		panic(err)
	}
}
