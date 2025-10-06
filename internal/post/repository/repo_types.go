package repository

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Post
type CreateOptions struct {
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

type ListOptions struct {
	Filter
}

type GetOptions struct {
	Filter
	PagQuery paginator.PaginatorQuery
}

type UpdateOptions struct {
	Post         models.Post
	Content      string
	FileIDs      []string
	Pin          bool
	TaggedTarget []string
}
