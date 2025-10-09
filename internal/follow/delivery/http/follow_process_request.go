package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/models"
	pkgErrors "github.com/hoag/go-social-feed/pkg/errors"
	"github.com/hoag/go-social-feed/pkg/jwt"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h handler) processCreateRequest(c *gin.Context) (createReq, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "follow.delivery.http.processDetailRequest.GetPayloadFromContext: unauthorized")
		return createReq{}, models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "follow.delivery.http.processCreateRequest.ShouldBindJSON: %v", err)
		return createReq{}, models.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "follow.delivery.http.processCreateRequest.Validate: %v", err)
		return createReq{}, models.Scope{}, errWrongBody
	}

	sc := jwt.NewScope(payload)

	return req, sc, nil
}

func (h handler) processDetailRequest(c *gin.Context) (string, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "follow.delivery.http.processDetailRequest.GetPayloadFromContext: unauthorized")
		return "", models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	id := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		h.l.Errorf(ctx, "follow.delivery.http.processDetailRequest.ObjectIDFromHex: %v", err)
		return "", models.Scope{}, errWrongBody
	}

	sc := jwt.NewScope(payload)

	return id, sc, nil
}

func (h handler) processGetRequest(c *gin.Context) (getReq, paginator.PaginatorQuery, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "follow.delivery.http.processGetRequest.GetPayloadFromContext: unauthorized")
		return getReq{}, paginator.PaginatorQuery{}, models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req getReq
	if err := c.ShouldBindQuery(&req); err != nil {
		h.l.Errorf(ctx, "follow.delivery.http.processGetRequest.ShouldBindQuery: %v", err)
		return getReq{}, paginator.PaginatorQuery{}, models.Scope{}, errWrongQuery
	}

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "follow.delivery.http.processGetRequest.Validate: %v", err)
		return getReq{}, paginator.PaginatorQuery{}, models.Scope{}, errWrongQuery
	}

	var pq paginator.PaginatorQuery
	if err := c.ShouldBindQuery(&pq); err != nil {
		h.l.Errorf(ctx, "follow.delivery.http.processGetRequest.ShouldBindQuery: %v", errWrongQuery)
		return getReq{}, paginator.PaginatorQuery{}, models.Scope{}, errWrongQuery
	}

	sc := jwt.NewScope(payload)

	return req, pq, sc, nil
}

func (h handler) processDeleteRequest(c *gin.Context) (string, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Errorf(ctx, "follow.delivery.http.processDeleteRequest.GetPayloadFromContext: unauthorized")
		return "", models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	id := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		h.l.Errorf(ctx, "follow.delivery.http.processDeleteRequest.ObjectIDFromHex: %v", err)
		return "", models.Scope{}, errWrongBody
	}

	sc := jwt.NewScope(payload)

	return id, sc, nil
}
