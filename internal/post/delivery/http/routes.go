package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/middleware"
)

func MapRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	// Public routes (no auth required)
	r.GET("/:id", h.Detail)
	r.GET("", h.Get)

	// Protected routes (auth required)
	authenticated := r.Group("")
	authenticated.Use(mw.Auth())
	authenticated.POST("", h.Create)
	authenticated.PUT("", h.Update)
	authenticated.DELETE("/:id", h.Delete)

	mapReactionRoutes(r, h)
}

func mapReactionRoutes(r *gin.RouterGroup, h Handler) {
	reaction := r.Group("/reaction")
	reaction.POST("", h.CreateReaction)
	reaction.GET("/:id", h.DetailReaction)
	reaction.GET("", h.GetReaction)
	reaction.DELETE("/:id", h.DeleteReaction)
}
