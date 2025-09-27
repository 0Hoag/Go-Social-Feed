package repository

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

//go:generate mockery --name=Repository
type Repository interface {
	PostRepo
	PostReactionRepo
}

type PostRepo interface {
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Post, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.Post, error)
	List(ctx context.Context, sc models.Scope, opts ListOptions) ([]models.Post, error)
	Get(ctx context.Context, sc models.Scope, opts GetOptions) ([]models.Post, paginator.PaginatorQuery, error)
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) error
	Delete(ctx context.Context, sc models.Scope, id string) error
}

type PostReactionRepo interface {
}
