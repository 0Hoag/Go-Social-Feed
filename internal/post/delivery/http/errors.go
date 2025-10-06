package http

import (
	"github.com/hoag/go-social-feed/internal/post"
	pkgErrors "github.com/hoag/go-social-feed/pkg/errors"
)

var (
	errWrongPaginationQuery = pkgErrors.NewHTTPError(140001, "Wrong pagination query")
	errWrongQuery           = pkgErrors.NewHTTPError(140002, "Wrong query")
	errWrongBody            = pkgErrors.NewHTTPError(140003, "Wrong body")

	// Post errors
	errPostVersionNotFound = pkgErrors.NewHTTPError(143004, "Post version not found")
	errPostNotFound        = pkgErrors.NewHTTPError(143005, "Post not found")
	errEmotionNotFound     = pkgErrors.NewHTTPError(143006, "Emotion not found")
	errPostEmotionNotFound = pkgErrors.NewHTTPError(143007, "Post emotion not found")
	errEmotionExists       = pkgErrors.NewHTTPError(143008, "Emotion exists")
)

func (h handler) mapError(err error) error {
	switch err {
	case post.ErrPostNotFound:
		return errPostNotFound
	case post.ErrPostVersionNotFound:
		return errPostVersionNotFound
	case post.ErrEmotionNotFound:
		return errEmotionNotFound
	case post.ErrPostEmotionNotFound:
		return errPostEmotionNotFound
	case post.ErrEmotionExists:
		return errEmotionExists
	default:
		panic(err)
	}
}
