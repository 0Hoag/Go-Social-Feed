package http

import (
	"github.com/hoag/go-social-feed/internal/scanner"
	pkgErrors "github.com/hoag/go-social-feed/pkg/errors"
)

var (
	errWrongBody = pkgErrors.NewHTTPError(140003, "Wrong body")
)

func (h handler) mapError(err error) error {
	switch err {
	case scanner.ErrScanToken:
		return errWrongBody
	default:
		panic(err)
	}
}
