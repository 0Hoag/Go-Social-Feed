package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/reaction"
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
	uc reaction.UseCase
}

func New(l log.Logger, uc reaction.UseCase) Handler {
	return handler{
		l:  l,
		uc: uc,
	}
}
