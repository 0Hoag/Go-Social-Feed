package http

import (
	"github.com/hoag/go-social-feed/internal/comment"
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"github.com/hoag/go-social-feed/pkg/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type attachments struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

func (r attachments) toInput() models.Attachment {
	return models.Attachment{
		Type: r.Type,
		URL:  r.Url,
	}
}

type createReq struct {
	PostID  string        `json:"post_id"`
	Content string        `json:"content"`
	Attach  []attachments `json:"attachments"`
}

func (r createReq) toInput() comment.CreateInput {
	var attach []models.Attachment
	for _, at := range r.Attach {
		attach = append(attach, at.toInput())
	}

	return comment.CreateInput{
		PostID:  r.PostID,
		Content: r.Content,
		Attach:  attach,
	}
}

func (r createReq) validate() error {
	if _, err := primitive.ObjectIDFromHex(r.PostID); err != nil {
		return errWrongBody
	}

	return nil
}

type updateReq struct {
	PostID  string        `json:"post_id"`
	Content string        `json:"content"`
	Attach  []attachments `json:"attachments"`
}

func (r updateReq) toInput() comment.UpdateInput {
	var attach []models.Attachment
	for _, at := range r.Attach {
		attach = append(attach, at.toInput())
	}

	return comment.UpdateInput{
		PostID:  r.PostID,
		Content: r.Content,
		Attach:  attach,
	}
}

func (r updateReq) validate() error {
	if _, err := primitive.ObjectIDFromHex(r.PostID); err != nil {
		return errWrongBody
	}

	return nil
}

type getReq struct {
	ID  string   `form:"id"`
	IDs []string `form:"ids[]"`
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

	return nil
}

func (r getReq) toFilter() comment.Filter {
	return comment.Filter{
		ID:  r.ID,
		IDs: r.IDs,
	}
}

func (h handler) newCommentDataResp(r models.Comment) commentDataResp {
	var attach []attachments
	for _, at := range r.Attachments {
		attach = append(attach, attachments{
			Type: at.Type,
			Url:  at.URL,
		})
	}

	return commentDataResp{
		ID:          r.ID.Hex(),
		PostID:      r.PostID.Hex(),
		AuthorID:    r.AuthorID.Hex(),
		Content:     r.Content,
		Attachments: attach,
		CreatedAt:   response.DateTime(r.CreatedAt),
	}
}

type detailResp struct {
	commentDataResp
}

func (h handler) newDetailResp(m models.Comment) detailResp {
	return detailResp{
		commentDataResp: h.newCommentDataResp(m),
	}
}

type commentDataResp struct {
	ID          string            `json:"id"`
	PostID      string            `json:"post_id"`
	AuthorID    string            `json:"author_id"`
	Content     string            `json:"content"`
	Attachments []attachments     `json:"attachments,omitempty"`
	CreatedAt   response.DateTime `json:"created_at"`
}

type commentItem struct {
	commentDataResp
}

type getMetaResponse struct {
	paginator.PaginatorResponse
}

type getResp struct {
	Items []commentItem   `json:"items"`
	Meta  getMetaResponse `json:"meta"`
}

func (h handler) newGetResp(out comment.GetOutput) getResp {
	items := make([]commentItem, 0, len(out.Comments))

	for _, p := range out.Comments {
		item := commentItem{
			commentDataResp: h.newCommentDataResp(p),
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
