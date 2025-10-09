package repository

import (
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Post
type CreateOptions struct {
	FollowerID string
	FolloweeID string
}

type Filter struct {
	ID         string
	IDs        []string
	FollowerID string
	FolloweeID string
}

type ListOptions struct {
	Filter
}

type GetOptions struct {
	Filter
	PagQuery paginator.PaginatorQuery
}
