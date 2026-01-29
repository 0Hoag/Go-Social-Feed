package post

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
)

//go:generate mockery --name=Usecase
type UseCase interface {
	PostUC
	Consumer
	ReactionUC
}

type PostUC interface {
	Create(ctx context.Context, sc models.Scope, input CreateInput) (models.Post, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.Post, error)
	List(ctx context.Context, sc models.Scope, input ListInput) ([]models.Post, error)
	Get(ctx context.Context, sc models.Scope, input GetInput) (GetOutput, error)
	GetOne(ctx context.Context, sc models.Scope, input GetOneInput) (models.Post, error)
	Update(ctx context.Context, sc models.Scope, input UpdateInput) error
	Delete(ctx context.Context, sc models.Scope, id string) error
}

type ReactionUC interface {
	CreateReaction(ctx context.Context, sc models.Scope, input CreateReactionInput) (models.Reaction, error)
	DetailReaction(ctx context.Context, sc models.Scope, id string) (models.Reaction, error)
	ListReaction(ctx context.Context, sc models.Scope, input ListReactionInput) ([]models.Reaction, error)
	GetReaction(ctx context.Context, sc models.Scope, input GetReactionInput) (GetReactionOutput, error)
	DeleteReaction(ctx context.Context, sc models.Scope, id string) error
}

type CommentUC interface {
	CreateComment(ctx context.Context, sc models.Scope, input CreateCommentInput) (models.Comment, error)
	DetailComment(ctx context.Context, sc models.Scope, id string) (models.Comment, error)
	ListComment(ctx context.Context, sc models.Scope, input ListCommentInput) ([]models.Comment, error)
	GetComment(ctx context.Context, sc models.Scope, input GetCommentInput) (GetOutput, error)
	UpdateComment(ctx context.Context, sc models.Scope, input UpdateCommentInput) (models.Comment, error)
	DeleteComment(ctx context.Context, sc models.Scope, id string) error
}

type Consumer interface {
	ProcessDeleteCommentMsg(ctx context.Context, sc models.Scope, input DeleteCommentMsgInput) error
	ProcessDeleteReactionMsg(ctx context.Context, sc models.Scope, input DeleteReactionMsgInput) error
}
