package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/models"
	pkgErrors "github.com/hoag/go-social-feed/pkg/errors"
	"github.com/hoag/go-social-feed/pkg/jwt"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) processCreateReactionRequest(c *gin.Context) (createReactionReq, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "post.delivery.http.processCreateReactionRequest.GetPayloadFromContext: unauthorized")
		return createReactionReq{}, models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req createReactionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "post.delivery.http.processCreateReactionRequest.ShouldBindJSON: %v", err)
		return createReactionReq{}, models.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "post.delivery.http.processCreateReactionRequest.Validate: %v", err)
		return createReactionReq{}, models.Scope{}, errWrongBody
	}

	sc := jwt.NewScope(payload)

	return req, sc, nil
}

func (h handler) processDetailReactionRequest(c *gin.Context) (string, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "post.delivery.http.processDetailReactionRequest.GetPayloadFromContext: unauthorized")
		return "", models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	id := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		h.l.Errorf(ctx, "post.delivery.http.processDetailReactionRequest.ObjectIDFromHex: %v", err)
		return "", models.Scope{}, errWrongBody
	}

	sc := jwt.NewScope(payload)

	return id, sc, nil
}

func (h handler) processGetReactionRequest(c *gin.Context) (getReactionReq, paginator.PaginatorQuery, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "reaction.delivery.http.processGetReactionRequest.GetPayloadFromContext: unauthorized")
		return getReactionReq{}, paginator.PaginatorQuery{}, models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req getReactionReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Errorf(ctx, "reaction.delivery.http.processGetReactionRequest.ShouldBindQuery: %v", err)
		return getReactionReq{}, paginator.PaginatorQuery{}, models.Scope{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "reaction.delivery.http.processGetReactionRequest.Validate: %v", err)
		return getReactionReq{}, paginator.PaginatorQuery{}, models.Scope{}, errWrongQuery
	}

	var pq paginator.PaginatorQuery
	if err := c.ShouldBindQuery(&pq); err != nil {
		h.l.Errorf(ctx, "reaction.delivery.http.processGetReactionRequest.ShouldBindQuery: %v", errWrongQuery)
		return getReactionReq{}, paginator.PaginatorQuery{}, models.Scope{}, errWrongQuery
	}

	sc := jwt.NewScope(payload)

	return req, pq, sc, nil
}

func (h handler) processDeleteReactionRequest(c *gin.Context) (string, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "post.delivery.http.processDeleteReactionRequest.GetPayloadFromContext: unauthorized")
		return "", models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	id := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		h.l.Errorf(ctx, "post.delivery.http.processDeleteReactionRequest.ObjectIDFromHex: %v", err)
		return "", models.Scope{}, errWrongBody
	}

	sc := jwt.NewScope(payload)

	return id, sc, nil
}
