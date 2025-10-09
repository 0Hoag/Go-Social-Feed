package repository

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Post
type CreateOptions struct {
	PostID string
	Type   models.ReactionType
}

type Filter struct {
	ID     string
	IDs    []string
	UserID string
	Type   models.ReactionType
}

type ListOptions struct {
	Filter
}

type GetOptions struct {
	Filter
	PagQuery paginator.PaginatorQuery
}
