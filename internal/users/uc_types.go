package users

import (
	"time"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Post
type CreateInput struct {
	UserName     string
	AvatarURL    string
	Phone        string
	PasswordHash string
	Birthday     time.Time
	Roles        []string
	Permissions  []string
}

type Filter struct {
	ID       string
	IDs      []string
	UserName string
	Phone    string
}

type GetOneInput struct {
	Filter
}

type ListInput struct {
	Filter
}

type GetInput struct {
	Filter
	PagQuery paginator.PaginatorQuery
}

type GetOutput struct {
	Users     []models.User
	Paginator paginator.Paginator
}

type UpdateInput struct {
	ID        string
	UserName  string
	AvatarURL string
}
