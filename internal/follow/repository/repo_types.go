package repository

import (
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Post
type CreateOptions struct {
	FolloweeID string
}

type Filter struct {
	ID         string
	IDs        []string
	AuthorID   string
	FolloweeID string
}

type ListOptions struct {
	Filter
}

type GetOptions struct {
	Filter
	PagQuery paginator.PaginatorQuery
}
