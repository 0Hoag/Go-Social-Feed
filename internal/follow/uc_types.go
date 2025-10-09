package follow

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Follow
type CreateInput struct {
	FollowerID string
	FolloweeID string
}

type Filter struct {
	ID         string
	IDs        []string
	FollowerID string
	FolloweeID string
}

type ListInput struct {
	Filter
}

type GetInput struct {
	Filter
	PagQuery paginator.PaginatorQuery
}

type GetOutput struct {
	Follows   []models.Follow
	Paginator paginator.Paginator
}
