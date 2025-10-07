package http

import (
	"github.com/gin-gonic/gin"
)

func (h handler) processLoginRequest(c *gin.Context) (loginReq, error) {
	ctx := c.Request.Context()

	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "post.delivery.http.processLoginRequest.ShouldBindJSON: %v", err)
		return loginReq{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "post.delivery.http.processLoginRequest.Validate: %v", err)
		return loginReq{}, errWrongBody
	}

	return req, nil
}
