package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/pkg/locale"
)

func (m Middleware) Locale() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := locale.GetLanguage(c)

		ctx := c.Request.Context()
		ctx = locale.SetLocaleToContext(ctx, l)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
