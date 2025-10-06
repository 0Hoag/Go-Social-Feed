package middleware

import (
	"github.com/hoag/go-social-feed/internal/users"
	"github.com/hoag/go-social-feed/pkg/log"

	pkgCrt "github.com/hoag/go-social-feed/pkg/encrypter"

	"github.com/hoag/go-social-feed/pkg/jwt"
)

type Middleware struct {
	l           log.Logger
	userUC      users.UseCase
	jwtManager  jwt.Manager
	encrypter   pkgCrt.Encrypter
	internalKey string
}

func New(
	l log.Logger,
	userUC users.UseCase,
	jwtManager jwt.Manager,
	encrypter pkgCrt.Encrypter,
	internalKey string,
) Middleware {
	return Middleware{
		l:           l,
		userUC:      userUC,
		jwtManager:  jwtManager,
		encrypter:   encrypter,
		internalKey: internalKey,
	}
}
