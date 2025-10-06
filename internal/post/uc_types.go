package post

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Post
type CreateInput struct {
	Pin          bool
	Content      string
	FileIDs      []string
	TaggedTarget []string
	Permission   string
}

type Filter struct {
	ID       string
	IDs      []string
	Pin      bool
	AuthorID string
}

type ListInput struct {
	Filter
}

type GetInput struct {
	Filter
	PagQuery paginator.PaginatorQuery
}

type GetOutput struct {
	Posts     []models.Post
	Paginator paginator.Paginator
}

type UpdateInput struct {
	ID           string
	Pin          bool
	Content      string
	FileIDs      []string
	TaggedTarget []string
	Permission   string
}
