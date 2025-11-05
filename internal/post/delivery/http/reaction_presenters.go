package http

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/pkg/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createReactionReq struct {
	PostID string `json:"post_id"`
	Type   string `json:"type"`
}

func (r createReactionReq) toInput() post.CreateReactionInput {
	return post.CreateReactionInput{
		PostID: r.PostID,
		Type:   models.ReactionType(r.Type),
	}
}

func (r createReactionReq) validate() error {
	if _, err := primitive.ObjectIDFromHex(r.PostID); err != nil {
		return errWrongBody
	}

	return nil
}

type getReactionReq struct {
	ID     string   `form:"id"`
	IDs    []string `form:"ids[]"`
	UserID string   `form:"user_id"`
	Type   string   `form:"type"`
}

func (r getReactionReq) validate() error {
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

	if r.UserID != "" {
		if _, err := primitive.ObjectIDFromHex(r.UserID); err != nil {
			return errWrongQuery
		}
	}

	return nil
}

func (r getReactionReq) toFilter() post.FilterReaction {
	return post.FilterReaction{
		ID:     r.ID,
		IDs:    r.IDs,
		UserID: r.UserID,
		Type:   models.ReactionType(r.Type),
	}
}

func (h handler) newReactionDataResp(r models.Reaction) reactionDataResp {
	return reactionDataResp{
		ID:        r.ID.Hex(),
		AuthorID:  r.AuthorID.Hex(),
		PostID:    r.PostID.Hex(),
		Type:      string(r.Type),
		CreatedAt: response.DateTime(r.CreatedAt),
	}
}

type detailReactionResp struct {
	reactionDataResp
}

func (h handler) newDetailReactionResp(p models.Reaction) detailReactionResp {
	return detailReactionResp{
		reactionDataResp: h.newReactionDataResp(p),
	}
}

type reactionDataResp struct {
	ID        string            `json:"id"`
	AuthorID  string            `json:"author_id"`
	PostID    string            `json:"post_id"`
	Type      string            `json:"type"`
	CreatedAt response.DateTime `json:"created_at"`
	UpdatedAt response.DateTime `json:"updated_at"`
}

type reactionItem struct {
	reactionDataResp
}

type getReactionResp struct {
	Items []reactionItem  `json:"items"`
	Meta  getMetaResponse `json:"meta"`
}

func (h handler) newGetReactionResp(out post.GetReactionOutput) getReactionResp {
	items := make([]reactionItem, 0, len(out.Reactions))

	for _, p := range out.Reactions {
		item := reactionItem{
			reactionDataResp: h.newReactionDataResp(p),
		}

		items = append(items, item)
	}

	return getReactionResp{
		Items: items,
		Meta: getMetaResponse{
			PaginatorResponse: out.Paginator.ToResponse(),
		},
	}
}
