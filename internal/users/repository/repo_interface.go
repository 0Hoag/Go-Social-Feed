package repository

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

//go:generate mockery --name=Repository
type Repository interface {
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.User, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.User, error)
	List(ctx context.Context, sc models.Scope, opts ListOptions) ([]models.User, error)
	Get(ctx context.Context, sc models.Scope, opts GetOptions) ([]models.User, paginator.Paginator, error)
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) error
	Delete(ctx context.Context, sc models.Scope, id string) error
}
