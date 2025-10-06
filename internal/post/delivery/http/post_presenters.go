package http

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"github.com/hoag/go-social-feed/pkg/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createReq struct {
	Pin          bool     `json:"pin"`
	Content      string   `json:"content"`
	FileIDs      []string `json:"file_ids"`
	TaggedTarget []string `json:"tagged_target"`
}

func (r createReq) toInput() post.CreateInput {
	return post.CreateInput{
		Pin:          r.Pin,
		Content:      r.Content,
		FileIDs:      r.FileIDs,
		TaggedTarget: r.TaggedTarget,
	}
}

func (r createReq) validate() error {
	// Validate that at least one of content, file_ids or share_post_id is provided
	if r.Content == "" && len(r.FileIDs) == 0 {
		return errWrongBody
	}

	if len(r.TaggedTarget) > 0 {
		for _, id := range r.TaggedTarget {
			if _, err := primitive.ObjectIDFromHex(id); err != nil {
				return errWrongBody
			}
		}
	}

	if len(r.FileIDs) > 0 {
		for _, id := range r.FileIDs {
			if _, err := primitive.ObjectIDFromHex(id); err != nil {
				return errWrongBody
			}
		}
	}

	return nil
}

type getReq struct {
	ID  string   `form:"id"`
	IDs []string `form:"ids[]"`
	Pin *bool    `form:"pin"`
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

func (r getReq) toFilter() post.Filter {
	return post.Filter{
		ID:  r.ID,
		IDs: r.IDs,
		Pin: *r.Pin,
	}
}

type updateReq struct {
	ID           string   `uri:"id"`
	Content      string   `json:"content"`
	FileIDs      []string `json:"file_ids"`
	TaggedTarget []string `json:"tagged_target"`
}

func (r updateReq) toInput() post.UpdateInput {
	var taggedTarget []string
	if len(r.TaggedTarget) > 0 {
		taggedTarget = r.TaggedTarget
	}

	return post.UpdateInput{
		ID:           r.ID,
		Content:      r.Content,
		FileIDs:      r.FileIDs,
		TaggedTarget: taggedTarget,
	}
}

func (r updateReq) validate() error {
	if _, err := primitive.ObjectIDFromHex(r.ID); err != nil {
		return errWrongBody
	}

	// Validate that at least one of content or file_ids is provided
	if r.Content == "" && len(r.FileIDs) == 0 {
		return errWrongBody
	}

	idArrays := [][]string{
		r.FileIDs,
		r.TaggedTarget,
	}

	// Add TaggedTarget arrays if not nil
	if len(r.TaggedTarget) > 0 {
		idArrays = append(idArrays, r.TaggedTarget)
	}

	for _, ids := range idArrays {
		for _, id := range ids {
			if _, err := primitive.ObjectIDFromHex(id); err != nil {
				return errWrongBody
			}
		}
	}

	return nil
}

func (h handler) newPostDataResp(p models.Post) postDataResp {
	return postDataResp{
		ID:           p.ID.Hex(),
		Content:      p.Content,
		Pin:          p.Pin,
		FileIDs:      util.ObjectIDsToHex(p.FileIDs),
		TaggedTarget: util.ObjectIDsToHex(p.TaggedTarget),
	}
}

type detailResp struct {
	postDataResp
}

func (h handler) newDetailResp(p models.Post) detailResp {
	return detailResp{
		postDataResp: h.newPostDataResp(p),
	}
}

type postDataResp struct {
	ID           string   `json:"id"`
	Content      string   `json:"content"`
	FileIDs      []string `json:"file_ids"`
	TaggedTarget []string `json:"tagged_target"`
	Pin          bool     `json:"pin"`
}

type postItem struct {
	postDataResp
}

type getMetaResponse struct {
	paginator.PaginatorResponse
}

type getResp struct {
	Items []postItem      `json:"items"`
	Meta  getMetaResponse `json:"meta"`
}

func (h handler) newGetResp(out post.GetOutput) getResp {
	items := make([]postItem, 0, len(out.Posts))

	for _, p := range out.Posts {
		item := postItem{
			postDataResp: h.newPostDataResp(p),
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
