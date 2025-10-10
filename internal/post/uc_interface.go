package post

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
)

//go:generate mockery --name=Usecase
type UseCase interface {
	PostUC
	Consumer
}

type PostUC interface {
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Post, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.Post, error)
	List(ctx context.Context, sc models.Scope, input ListInput) ([]models.Post, error)
	Get(ctx context.Context, sc models.Scope, input GetInput) (GetOutput, error)
	Update(ctx context.Context, sc models.Scope, input UpdateInput) error
	Delete(ctx context.Context, sc models.Scope, id string) error
}

type Consumer interface {
	ProcessDeleteCommentMsg(ctx context.Context, sc models.Scope, input DeleteCommentMsgInput) error
	ProcessDeleteReactionMsg(ctx context.Context, sc models.Scope, input DeleteReactionMsgInput) error
}
