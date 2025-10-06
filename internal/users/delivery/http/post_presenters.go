package http

import (
	"time"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/users"
	"github.com/hoag/go-social-feed/pkg/paginator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createReq struct {
	Username string    `json:"username"`
	Phone    string    `json:"phone"`
	Password string    `json:"password"`
	Birthday time.Time `json:"birthday"`
}

func (r createReq) toInput() users.CreateInput {
	return users.CreateInput{
		UserName:     r.Username,
		Phone:        r.Phone,
		PasswordHash: r.Password,
		Birthday:     r.Birthday,
	}
}

func (r createReq) validate() error {

	return nil
}

type getReq struct {
	ID       string   `form:"id"`
	IDs      []string `form:"ids[]"`
	Username string   `form:"username"`
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

func (r getReq) toFilter() users.Filter {
	return users.Filter{
		ID:       r.ID,
		IDs:      r.IDs,
		UserName: r.Username,
	}
}

type updateReq struct {
	ID        string `uri:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

func (r updateReq) toInput() users.UpdateInput {
	return users.UpdateInput{
		ID:        r.ID,
		UserName:  r.Username,
		AvatarURL: r.AvatarURL,
	}
}

func (r updateReq) validate() error {
	if _, err := primitive.ObjectIDFromHex(r.ID); err != nil {
		return errWrongBody
	}

	return nil
}

func (h handler) newusersDataResp(p models.User) usersDataResp {
	return usersDataResp{
		ID: p.ID.Hex(),
	}
}

type detailResp struct {
	usersDataResp
}

func (h handler) newDetailResp(p models.User) detailResp {
	return detailResp{
		usersDataResp: h.newusersDataResp(p),
	}
}

type usersDataResp struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatar_url"`
}

type usersItem struct {
	usersDataResp
}

type getMetaResponse struct {
	paginator.PaginatorResponse
}

type getResp struct {
	Items []usersItem     `json:"items"`
	Meta  getMetaResponse `json:"meta"`
}

func (h handler) newGetResp(out users.GetOutput) getResp {
	items := make([]usersItem, 0, len(out.Users))

	for _, p := range out.Users {
		item := usersItem{
			usersDataResp: h.newusersDataResp(p),
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
