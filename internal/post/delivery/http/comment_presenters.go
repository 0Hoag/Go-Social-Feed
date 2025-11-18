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

type createCommentReq struct {
	PostID  string        `json:"post_id"`
	Content string        `json:"content"`
	Attach  []attachments `json:"attachments"`
}

func (r createCommentReq) toInput() comment.CreateInput {
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

func (r createCommentReq) validate() error {
	if _, err := primitive.ObjectIDFromHex(r.PostID); err != nil {
		return errWrongBody
	}

	return nil
}

type updateCommentReq struct {
	PostID  string        `json:"post_id"`
	Content string        `json:"content"`
	Attach  []attachments `json:"attachments"`
}

func (r updateCommentReq) toInput() comment.UpdateInput {
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

func (r updateCommentReq) validate() error {
	if _, err := primitive.ObjectIDFromHex(r.PostID); err != nil {
		return errWrongBody
	}

	return nil
}

type getCommentReq struct {
	ID  string   `form:"id"`
	IDs []string `form:"ids[]"`
}

func (r getCommentReq) validate() error {
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

func (r getCommentReq) toFilter() comment.Filter {
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

type detailCommentResp struct {
	commentDataResp
}

func (h handler) newCommentDetailResp(m models.Comment) detailCommentResp {
	return detailCommentResp{
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

type getCommentMetaResponse struct {
	paginator.PaginatorResponse
}

type getCommentResp struct {
	Items []commentItem   `json:"items"`
	Meta  getMetaResponse `json:"meta"`
}

func (h handler) newGetCommentResp(out comment.GetOutput) getCommentResp {
	items := make([]commentItem, 0, len(out.Comments))

	for _, p := range out.Comments {
		item := commentItem{
			commentDataResp: h.newCommentDataResp(p),
		}

		items = append(items, item)
	}

	return getCommentResp{
		Items: items,
		Meta: getMetaResponse{
			PaginatorResponse: out.Paginator.ToResponse(),
		},
	}
}
