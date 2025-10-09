package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/middleware"
)

func MapRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {

	reactionGroup := r.Group("")
	reactionGroup.Use(mw.Auth())
	{
		reactionGroup.POST("", h.Create)
		reactionGroup.GET("/:id", h.Detail)
		reactionGroup.GET("", h.Get)
		reactionGroup.DELETE("/:id", h.Delete)
	}
}
