package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hoag/go-social-feed/internal/users"
	"github.com/hoag/go-social-feed/pkg/response"
)

// @Summary Create user
// @Schemes
// @Description Create user
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
// @Router /news-feed/user [POST]
func (h handler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	req, err := h.processCreateRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "user.delivery.http.Create.processCreateRequest: %v", err)
		response.Error(c, err)
		return
	}

	e, err := h.uc.Create(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "user.delivery.http.Create.Create: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailResp(e))
}

// @Summary Get user detail
// @Schemes
// @Description Get user detail
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
// @Router /news-feed/user/{id} [GET]
func (h handler) Detail(c *gin.Context) {
	ctx := c.Request.Context()

	id, sc, err := h.processDetailRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "user.delivery.http.Detail.processDetailRequest: %v", err)
		response.Error(c, err)
		return
	}

	p, err := h.uc.Detail(ctx, sc, id)
	if err != nil {
		h.l.Errorf(ctx, "user.delivery.http.Detail.Detail: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailResp(p))
}

// @Summary Get user
// @Schemes
// @Description Get user with filter
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
// @Router /news-feed/user [GET]
func (h handler) Get(c *gin.Context) {
	ctx := c.Request.Context()

	req, paq, sc, err := h.processGetRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "user.delivery.http.Get.processGetRequest: %v", err)
		response.Error(c, err)
		return
	}

	var input users.GetInput
	input.Filter = req.toFilter()
	paq.Adjust()
	input.PagQuery = paq

	e, err := h.uc.Get(ctx, sc, input)
	if err != nil {
		h.l.Errorf(ctx, "user.delivery.http.Get.Get: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newGetResp(e))
}

// @Summary Update user
// @Schemes
// @Description Update user
// @Param Access-Control-Allow-Origin header string false "Access-Control-Allow-Origin" default(*)
// @Param User-Agent header string false "User-Agent" default(Swagger-Codegen/1.0.0/go)
// @Param Authorization header string true "Bearer JWT token" default(Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vcC50YW5jYS52bi9hcGkvdjQvYXV0aC9zaWduaW4tdjIiLCJpYXQiOjE3MTY1ODUyNDAsIm5iZiI6MTcxNjU4NTI0MCwianRpIjoidFBJMldUa0JldThYdnJMZiIsInN1YiI6Ik5pdEpwZUp1dkF4M1pjYXdKIiwicHJ2IjoiMWM1NTIwZjcwYmFhNjU1ZGRjNTc0NmE2NzY0ZjM3MmExYjY1NWFhNiIsInNob3BfaWQiOiI1YzIwYTE5YzBiMDg4ODBmNTk0ZmM0NjgiLCJzaG9wX3VzZXJuYW1lIjoicmF2ZSIsInNob3BfcHJlZml4IjoidCIsInR5cGUiOiJhcGkifQ.DnxirM5IXQY3B9Vcc6Qco7c9f0ABGjoeLu_1LfHiRjE)"
// @Param lang header string false "Language" default(en)
// @Param id path string true "ID"
// @Param body body updateReq true "Body"
// @Produce json
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} detailResp
// @Failure 400 {object} response.Resp "Bad Request"
// @Router /news-feed/user [PUT]
func (h handler) Update(c *gin.Context) {
	ctx := c.Request.Context()

	req, sc, err := h.processUpdateRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "user.delivery.http.Update.processUpdateRequest: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.Update(ctx, sc, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "user.delivery.http.Update.Update: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, nil)
}

// @Summary Delete user
// @Schemes
// @Description Delete user
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
// @Router /news-feed/user/{id} [DELETE]
func (h handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	id, sc, err := h.processDeleteRequest(c)
	if err != nil {
		h.l.Errorf(ctx, "user.delivery.http.Delete.processDeleteRequest: %v", err)
		response.Error(c, err)
		return
	}

	err = h.uc.Delete(ctx, sc, id)
	if err != nil {
		h.l.Errorf(ctx, "user.delivery.http.Delete.Delete: %v", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, nil)
}
