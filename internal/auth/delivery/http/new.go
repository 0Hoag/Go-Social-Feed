package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/auth"
	"github.com/hoag/go-social-feed/pkg/log"
)

type Handler interface {
	Login(c *gin.Context)
}

type handler struct {
	l  log.Logger
	uc auth.UseCase
}

func New(l log.Logger, uc auth.UseCase) Handler {
	return handler{
		l:  l,
		uc: uc,
	}
}
