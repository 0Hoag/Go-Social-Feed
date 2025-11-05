package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/middleware"
)

func MapRoutes(r *gin.RouterGroup, h Handler, mw middleware.Middleware) {
	r.Use(mw.Auth())
	mapPostRoutes(r, h)
	mapReactionRoutes(r, h)
}

func mapPostRoutes(r *gin.RouterGroup, h Handler) {
	r.POST("", h.Create)
	r.GET("/:id", h.Detail)
	r.GET("", h.Get)
	r.PUT("", h.Update)
	r.DELETE("/:id", h.Delete)
}

func mapReactionRoutes(r *gin.RouterGroup, h Handler) {
	reaction := r.Group("/reaction")
	reaction.POST("", h.CreateReaction)
	reaction.GET("/:id", h.DetailReaction)
	reaction.GET("", h.GetReaction)
	reaction.DELETE("/:id", h.DeleteReaction)
}
