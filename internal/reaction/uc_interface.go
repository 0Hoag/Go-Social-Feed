package reaction

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
)

//go:generate mockery --name=Usecase
type UseCase interface {
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Reaction, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.Reaction, error)
	List(ctx context.Context, sc models.Scope, input ListInput) ([]models.Reaction, error)
	Get(ctx context.Context, sc models.Scope, input GetInput) (GetOutput, error)
	Delete(ctx context.Context, sc models.Scope, id string) error
}
