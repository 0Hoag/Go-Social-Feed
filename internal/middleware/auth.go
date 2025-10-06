package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/pkg/jwt"
	"github.com/hoag/go-social-feed/pkg/response"
)

func (m Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimSpace(strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer "))
		if tokenString == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		payload, err := m.jwtManager.Verify(tokenString)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx = jwt.SetPayloadToContext(ctx, payload)

		// Set scope to context
		scope := jwt.NewScope(payload)
		ctx = jwt.SetScopeToContext(ctx, scope)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
