package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/middleware"
)

func MapRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.Use(mw.Auth())
	r.POST("", h.Create)
	r.GET("/:id", h.Detail)
	r.GET("", h.Get)
	r.DELETE("/:id", h.Delete)
}
