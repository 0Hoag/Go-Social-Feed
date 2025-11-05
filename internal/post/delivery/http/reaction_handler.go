package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/pkg/response"
)

// @Summary Create reaction
// @Schemes
// @Description Create reaction
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vcC50YW5jYS52bi9hcGkvdjQvYXV0aC9zaWduaW4tdjIiLCJpYXQiOjE3MTY1ODUyNDAsIm5iZiI6MTcxNjU4NTI0MCwianRpIjoidFBJMldUa0JldThYdnJMZiIsInN1YiI6Ik5pdEpwZUp1dkF4M1pjYXdKIiwicHJ2IjoiMWM1NTIwZjcwYmFhNjU1ZGRjNTc0NmE2NzY0ZjM3MmExYjY1NWFhNiIsInNob3BfaWQiOiI1YzIwYTE5YzBiMDg4ODBmNTk0ZmM0NjgiLCJzaG9wX3VzZXJuYW1lIjoicmF2ZSIsInNob3BfcHJlZml4IjoidCIsInR5cGUiOiJhcGkifQ.DnxirM5IXQY3B9Vcc6Qco7c9f0ABGjoeLu_1LfHiRjE)"
// @Param lang header string false "Language" default(en)
// @Param body body createReq true "Body"
// @Produce json
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} detailResp
// @Failure 400 {object} response.Resp "Bad Request"
// @Router /news-feed/post/reaction [POST]
func (h handler) CreateReaction(c *gin.Context) {
	ctx := c.Request.Context()

	req, sc, err := h.processCreateReactionRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "post.delivery.http.CreateReaction.processCreateReactionRequest: %v", err)
		response.Error(c, err)
		return
	}

	e, err := h.uc.CreateReaction(ctx, sc, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "post.delivery.http.CreateReaction.CreateReaction: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailReactionResp(e))
}

// @Summary Get reaction detail
// @Schemes
// @Description Get reaction detail
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vcC50YW5jYS52bi9hcGkvdjQvYXV0aC9zaWduaW4tdjIiLCJpYXQiOjE3MTY1ODUyNDAsIm5iZiI6MTcxNjU4NTI0MCwianRpIjoidFBJMldUa0JldThYdnJMZiIsInN1YiI6Ik5pdEpwZUp1dkF4M1pjYXdKIiwicHJ2IjoiMWM1NTIwZjcwYmFhNjU1ZGRjNTc0NmE2NzY0ZjM3MmExYjY1NWFhNiIsInNob3BfaWQiOiI1YzIwYTE5YzBiMDg4ODBmNTk0ZmM0NjgiLCJzaG9wX3VzZXJuYW1lIjoicmF2ZSIsInNob3BfcHJlZml4IjoidCIsInR5cGUiOiJhcGkifQ.DnxirM5IXQY3B9Vcc6Qco7c9f0ABGjoeLu_1LfHiRjE)"
// @Param lang header string false "Language" default(en)
// @Param id path string true "ID"
// @Produce json
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} detailResp
// @Failure 400 {object} response.Resp "Bad Request"
// @Router /news-feed/post/reaction/{id} [GET]
func (h handler) DetailReaction(c *gin.Context) {
	ctx := c.Request.Context()

	id, sc, err := h.processDetailReactionRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "post.delivery.http.DetailReaction.processDetailReactionRequest: %v", err)
		response.Error(c, err)
		return
	}

	p, err := h.uc.DetailReaction(ctx, sc, id)
	if err != nil {
		h.l.Errorf(ctx, "post.delivery.http.DetailReaction.DetailReaction: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailReactionResp(p))
}

// @Summary Get reaction
// @Schemes
// @Description Get reaction with filter
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vcC50YW5jYS52bi9hcGkvdjQvYXV0aC9zaWduaW4tdjIiLCJpYXQiOjE3MTY1ODUyNDAsIm5iZiI6MTcxNjU4NTI0MCwianRpIjoidFBJMldUa0JldThYdnJMZiIsInN1YiI6Ik5pdEpwZUp1dkF4M1pjYXdKIiwicHJvIjoiMWM1NTIwZjcwYmFhNjU1ZGRjNTc0NmE2NzY0ZjM3MmExYjY1NWFhNiIsInNob3BfaWQiOiI1YzIwYTE5YzBiMDg4ODBmNTk0ZmM0NjgiLCJzaG9wX3VzZXJuYW1lIjoicmF2ZSIsInNob3BfcHJlZml4IjoidCIsInR5cGUiOiJhcGkifQ.DnxirM5IXQY3B9Vcc6Qco7c9f0ABGjoeLu_1LfHiRjE)
// @Param lang header string false "Language" default(en)
// @Param id query string false "ID"
// @Param ids query []string false "IDs" collectionFormat(csv)
// @Param type query string false "Type" Enums(company, self)
// @Param status query string false "Status" Enums(approved, pending, rejected,pending_update)
// @Param pin query boolean false "Pin"
// @Param user_id query string false "User ID"
// @Produce json
// @Tags Users
// @Accept json
// @Success 200 {object} getResp
// @Failure 400 {object} response.Resp "Bad Request"
// @Router /news-feed/post/reaction [GET]
func (h handler) GetReaction(c *gin.Context) {
	ctx := c.Request.Context()

	req, paq, sc, err := h.processGetReactionRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "post.delivery.http.GetReaction.processGetRequest: %v", err)
		response.Error(c, err)
		return
	}

	var input post.GetReactionInput
	input.FilterReaction = req.toFilter()
	paq.Adjust()
	input.PagQuery = paq

	e, err := h.uc.GetReaction(ctx, sc, input)
	if err != nil {
		h.l.Errorf(ctx, "post.delivery.http.GetReaction.GetReaction: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newGetReactionResp(e))
}

// @Summary Delete reaction
// @Schemes
// @Description Delete reaction
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vcC50YW5jYS52bi9hcGkvdjQvYXV0aC9zaWduaW4tdjIiLCJpYXQiOjE3MTY1ODUyNDAsIm5iZiI6MTcxNjU4NTI0MCwianRpIjoidFBJMldUa0JldThYdnJMZiIsInN1YiI6Ik5pdEpwZUp1dkF4M1pjYXdKIiwicHJ2IjoiMWM1NTIwZjcwYmFhNjU1ZGRjNTc0NmE2NzY0ZjM3MmExYjY1NWFhNiIsInNob3BfaWQiOiI1YzIwYTE5YzBiMDg4ODBmNTk0ZmM0NjgiLCJzaG9wX3VzZXJuYW1lIjoicmF2ZSIsInNob3BfcHJlZml4IjoidCIsInR5cGUiOiJhcGkifQ.DnxirM5IXQY3B9Vcc6Qco7c9f0ABGjoeLu_1LfHiRjE)"
// @Param lang header string false "Language" default(en)
// @Param id path string true "ID"
// @Produce json
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} detailResp
// @Failure 400 {object} response.Resp "Bad Request"
// @Router /news-feed/post/reaction/{id} [DELETE]
func (h handler) DeleteReaction(c *gin.Context) {
	ctx := c.Request.Context()

	id, sc, err := h.processDeleteReactionRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "post.delivery.http.DeleteReaction.processDeleteRequest: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.DeleteReaction(ctx, sc, id)
	if err != nil {
		h.l.Errorf(ctx, "post.delivery.http.DeleteReaction.DeleteReaction: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, nil)
}
