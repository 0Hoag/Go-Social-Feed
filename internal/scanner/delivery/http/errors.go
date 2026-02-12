package http

import (
	"github.com/hoag/go-social-feed/internal/scanner"
	pkgErrors "github.com/hoag/go-social-feed/pkg/errors"
)

var (
	errWrongBody          = pkgErrors.NewHTTPError(140003, "Wrong body")
	errTokenNotFound      = pkgErrors.NewHTTPError(140404, "Token not found on DexScreener. Please check the token symbol or address.")
	errSourceCodeNotFound = pkgErrors.NewHTTPError(140405, "Source code not found on supported networks (ETH, BSC, BASE, ARBITRUM, POLYGON). This token may not be a smart contract or is on an unsupported network.")
)

func (h handler) mapError(err error) error {
	switch err {
	case scanner.ErrTokenNotFound:
		return errTokenNotFound
	case scanner.ErrSourceCodeNotFound:
		return errSourceCodeNotFound
	default:
		panic(err)
	}
}
