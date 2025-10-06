package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/pkg/jwt"
	"github.com/hoag/go-social-feed/pkg/response"
)

// CheckPermission validates permission locally using roles/claims in JWT payload
func (m Middleware) CheckPermission(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, ok := jwt.GetPayloadFromContext(c.Request.Context())
		if !ok {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		// Local permission check based on roles/permissions in payload
		has := false
		// Prefer explicit permissions field; fallback to roles contains required
		if len(payload.Permissions) > 0 {
			for _, p := range payload.Permissions {
				if p == required {
					has = true
					break
				}
			}
		} else if len(payload.Roles) > 0 {
			for _, r := range payload.Roles {
				if strings.EqualFold(r, required) {
					has = true
					break
				}
			}
		}

		if !has {
			response.Error(c, errPermission)
			c.Abort()
			return
		}

		c.Next()
	}
}
