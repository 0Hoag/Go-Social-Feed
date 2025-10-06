package middleware

import (
	"crypto/subtle"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/pkg/response"
)

func (m Middleware) AuthInternalKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		secKeyEncoded := strings.TrimSpace(c.GetHeader("Internal-Key"))
		if secKeyEncoded == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		secKey, err := m.encrypter.Decrypt(secKeyEncoded)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		if subtle.ConstantTimeCompare([]byte(secKey), []byte(m.internalKey)) != 1 {
			response.Unauthorized(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
