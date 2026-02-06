package http

import "github.com/gin-gonic/gin"

func (h handler) processRequest(c *gin.Context) (scannerTokenInput, error) {
	ctx := c.Request.Context()

	var req scannerTokenInput
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Errorf(ctx, "scanner.delivery.http.processRequest.ShouldBindQuery: %v", err)
		return scannerTokenInput{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "scanner.delivery.http.processRequest.Validate: %v", err)
		return scannerTokenInput{}, errWrongBody
	}

	return req, nil
}
