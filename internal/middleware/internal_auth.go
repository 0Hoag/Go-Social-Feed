package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/pkg/jwt"
	"github.com/hoag/go-social-feed/pkg/response"
)

func (m Middleware) AuthInternal() gin.HandlerFunc {
	return func(c *gin.Context) {
		scopeString := strings.TrimSpace(c.GetHeader("Scope"))
		if scopeString == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		scope, err := jwt.ParseScopeHeader(scopeString)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx = jwt.SetScopeToContext(ctx, scope)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
