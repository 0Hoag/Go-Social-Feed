package http

import (
	"github.com/hoag/go-social-feed/internal/follow"
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"github.com/hoag/go-social-feed/pkg/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createReq struct {
	FollowerID string `json:"follower_id"`
	FolloweeID string `json:"followee_id"`
}

func (r createReq) toInput() follow.CreateInput {
	return follow.CreateInput{
		FollowerID: r.FollowerID,
		FolloweeID: r.FolloweeID,
	}
}

func (r createReq) validate() error {
	if _, err := primitive.ObjectIDFromHex(r.FollowerID); err != nil {
		return errWrongBody
	}

	if _, err := primitive.ObjectIDFromHex(r.FolloweeID); err != nil {
		return errWrongBody
	}

	return nil
}

type getReq struct {
	ID         string   `form:"id"`
	IDs        []string `form:"ids[]"`
	FollowerID string   `form:"follower_id"`
	FolloweeID string   `form:"followee_id"`
}

func (r getReq) validate() error {
	if len(r.IDs) > 0 {
		for _, id := range r.IDs {
			if _, err := primitive.ObjectIDFromHex(id); err != nil {
				return errWrongQuery
			}
		}
	}

	if r.ID != "" {
		if _, err := primitive.ObjectIDFromHex(r.ID); err != nil {
			return errWrongQuery
		}
	}

	if r.FollowerID != "" {
		if _, err := primitive.ObjectIDFromHex(r.FollowerID); err != nil {
			return errWrongQuery
		}
	}

	if r.FolloweeID != "" {
		if _, err := primitive.ObjectIDFromHex(r.FolloweeID); err != nil {
			return errWrongQuery
		}
	}

	return nil
}

func (r getReq) toFilter() follow.Filter {
	return follow.Filter{
		ID:         r.ID,
		IDs:        r.IDs,
		FollowerID: r.FollowerID,
		FolloweeID: r.FolloweeID,
	}
}

func (h handler) newFollowDataResp(r models.Follow) followDataResp {
	return followDataResp{
		ID:         r.ID.Hex(),
		FollowerID: r.FollowerID.Hex(),
		FolloweeID: r.FolloweeID.Hex(),
		CreatedAt:  response.DateTime(r.CreatedAt),
	}
}

type detailResp struct {
	followDataResp
}

func (h handler) newDetailResp(m models.Follow) detailResp {
	return detailResp{
		followDataResp: h.newFollowDataResp(m),
	}
}

type followDataResp struct {
	ID         string            `json:"id"`
	FollowerID string            `json:"follower_id"`
	FolloweeID string            `json:"followee_id"`
	CreatedAt  response.DateTime `json:"created_at"`
}

type followItem struct {
	followDataResp
}

type getMetaResponse struct {
	paginator.PaginatorResponse
}

type getResp struct {
	Items []followItem    `json:"items"`
	Meta  getMetaResponse `json:"meta"`
}

func (h handler) newGetResp(out follow.GetOutput) getResp {
	items := make([]followItem, 0, len(out.Follows))

	for _, p := range out.Follows {
		item := followItem{
			followDataResp: h.newFollowDataResp(p),
		}

		items = append(items, item)
	}

	return getResp{
		Items: items,
		Meta: getMetaResponse{
			PaginatorResponse: out.Paginator.ToResponse(),
		},
	}
}
