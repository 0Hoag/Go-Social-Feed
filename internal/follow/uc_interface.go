package follow

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
)

//go:generate mockery --name=Usecase
type UseCase interface {
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Follow, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.Follow, error)
	List(ctx context.Context, sc models.Scope, input ListInput) ([]models.Follow, error)
	Get(ctx context.Context, sc models.Scope, input GetInput) (GetOutput, error)
	Delete(ctx context.Context, sc models.Scope, id string) error
}
