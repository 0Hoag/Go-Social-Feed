package comment

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
)

//go:generate mockery --name=Usecase
type UseCase interface {
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Comment, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.Comment, error)
	List(ctx context.Context, sc models.Scope, input ListInput) ([]models.Comment, error)
	Get(ctx context.Context, sc models.Scope, input GetInput) (GetOutput, error)
	Update(ctx context.Context, sc models.Scope, input UpdateInput) (models.Comment, error)
	Delete(ctx context.Context, sc models.Scope, id string) error
}
