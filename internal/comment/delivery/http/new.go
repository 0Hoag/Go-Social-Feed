package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/comment"
	"github.com/hoag/go-social-feed/pkg/log"
)

type Handler interface {
	Create(c *gin.Context)
	Detail(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type handler struct {
	l  log.Logger
	uc comment.UseCase
}

func New(l log.Logger, uc comment.UseCase) Handler {
	return handler{
		l:  l,
		uc: uc,
	}
}
