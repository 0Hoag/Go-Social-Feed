package reaction

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Post
type CreateInput struct {
	PostID string
	Type   models.ReactionType
}

type Filter struct {
	ID     string
	IDs    []string
	UserID string
	Type   models.ReactionType
}

type ListInput struct {
	Filter
}

type GetInput struct {
	Filter
	PagQuery paginator.PaginatorQuery
}

type GetOutput struct {
	Reactions []models.Reaction
	Paginator paginator.Paginator
}
