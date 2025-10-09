package repository

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Follow, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.Follow, error)
	List(ctx context.Context, sc models.Scope, opts ListOptions) ([]models.Follow, error)
	Get(ctx context.Context, sc models.Scope, opts GetOptions) ([]models.Follow, paginator.Paginator, error)
	Delete(ctx context.Context, sc models.Scope, id string) error
}
