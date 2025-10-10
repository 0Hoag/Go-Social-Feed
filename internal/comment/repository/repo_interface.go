package repository

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Comment, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.Comment, error)
	GetOne(ctx context.Context, sc models.Scope, opts GetOneOptions) (models.Comment, error)
	List(ctx context.Context, sc models.Scope, opts ListOptions) ([]models.Comment, error)
	Get(ctx context.Context, sc models.Scope, opts GetOptions) ([]models.Comment, paginator.Paginator, error)
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) (models.Comment, error)
	Delete(ctx context.Context, sc models.Scope, id string) error
}
