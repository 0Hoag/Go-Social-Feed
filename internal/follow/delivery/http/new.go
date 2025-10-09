package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/follow"
	"github.com/hoag/go-social-feed/pkg/log"
)

type Handler interface {
	Create(c *gin.Context)
	Detail(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
}

type handler struct {
	l  log.Logger
	uc follow.UseCase
}

func New(l log.Logger, uc follow.UseCase) Handler {
	return handler{
		l:  l,
		uc: uc,
	}
}
