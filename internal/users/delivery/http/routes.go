package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/middleware"
)

func MapRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.POST("", h.Create)

	authGroup := r.Group("")
	authGroup.Use(mw.Auth())
	{
		authGroup.GET("/:id", h.Detail)
		authGroup.GET("", h.Get)
		authGroup.PUT("", h.Update)
		authGroup.DELETE("/:id", h.Delete)
	}
}
