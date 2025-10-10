package comment

import (
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

// Comment
type CreateInput struct {
	PostID  string
	Content string
	Attach  []models.Attachment
}

type Filter struct {
	ID     string
	IDs    []string
	PostID string
}

type ListInput struct {
	Filter
}

type GetInput struct {
	Filter
	PagQuery paginator.PaginatorQuery
}

type GetOutput struct {
	Comments  []models.Comment
	Paginator paginator.Paginator
}

type UpdateInput struct {
	PostID  string
	Content string
	Attach  []models.Attachment
}
