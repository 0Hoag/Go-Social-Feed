package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/pkg/response"
)

// @Summary Scanner token
// @Schemes
// @Description Scanner token
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vcC50YW5jYS52bi9hcGkvdjQvYXV0aC9zaWduaW4tdjIiLCJpYXQiOjE3MTY1ODUyNDAsIm5iZiI6MTcxNjU4NTI0MCwianRpIjoidFBJMldUa0JldThYdnJMZiIsInN1YiI6Ik5pdEpwZUp1dkF4M1pjYXdKIiwicHJ2IjoiMWM1NTIwZjcwYmFhNjU1ZGRjNTc0NmE2NzY0ZjM3MmExYjY1NWFhNiIsInNob3BfaWQiOiI1YzIwYTE5YzBiMDg4ODBmNTk0ZmM0NjgiLCJzaG9wX3VzZXJuYW1lIjoicmF2ZSIsInNob3BfcHJlZml4IjoidCIsInR5cGUiOiJhcGkifQ.DnxirM5IXQY3B9Vcc6Qco7c9f0ABGjoeLu_1LfHiRjE)"
// @Param lang header string false "Language" default(en)
// @Param id query string false "token"
// @Tags Scanner
// @Accept json
// @Produce json
// @Success 200 {object} detailResp
// @Failure 400 {object} response.Resp "Bad Request"
// @Router /news-feed/scanner [GET]
func (h handler) ScanToken(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "scanner.delivery.http.ScanToken: %v", err)
		response.Error(c, err)
		return
	}

	token, err := h.uc.ScanToken(ctx, req.ToScanTokenInput())
	if err != nil {
		h.l.Errorf(ctx, "scanner.delivery.http.ScanToken: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.ToScanTokenOutput(token))
}
