package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/pkg/jwt"
	"github.com/hoag/go-social-feed/pkg/response"
)

func (m Middleware) UserSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		scope, ok := jwt.GetScopeFromContext(ctx)
		if !ok {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		sessionUser, err := m.userUC.GetSessionUser(ctx, scope)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		ctx = jwt.SetUserToContext(ctx, sessionUser)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
