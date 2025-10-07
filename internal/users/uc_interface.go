package users

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
)

//go:generate mockery --name=Usecase
type UseCase interface {
	Create(ctx context.Context, input CreateInput) (models.User, error)
	GetSessionUser(ctx context.Context, sc models.Scope) (models.User, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.User, error)
	GetOne(ctx context.Context, f Filter) (models.User, error)
	List(ctx context.Context, sc models.Scope, input ListInput) ([]models.User, error)
	Get(ctx context.Context, sc models.Scope, input GetInput) (GetOutput, error)
	Update(ctx context.Context, sc models.Scope, input UpdateInput) error
	Delete(ctx context.Context, sc models.Scope, id string) error
}
