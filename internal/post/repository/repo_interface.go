package repository

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/pkg/paginator"
)

//go:generate mockery --name=Repository
type Repository interface {
	PostRepo
	ReactionRepo
	CommentRepo
}

type PostRepo interface {
	Create(ctx context.Context, sc models.Scope, opts CreateOptions) (models.Post, error)
	Detail(ctx context.Context, sc models.Scope, id string) (models.Post, error)
	List(ctx context.Context, sc models.Scope, opts ListOptions) ([]models.Post, error)
	Get(ctx context.Context, sc models.Scope, opts GetOptions) ([]models.Post, paginator.Paginator, error)
	Update(ctx context.Context, sc models.Scope, opts UpdateOptions) error
	Delete(ctx context.Context, sc models.Scope, id string) error
}

type ReactionRepo interface {
	CreateReaction(ctx context.Context, sc models.Scope, opts CreateReactionOptions) (models.Reaction, error)
	DetailReaction(ctx context.Context, sc models.Scope, id string) (models.Reaction, error)
	ListReaction(ctx context.Context, sc models.Scope, opts ListReactionOptions) ([]models.Reaction, error)
	GetReaction(ctx context.Context, sc models.Scope, opts GetReactionOptions) ([]models.Reaction, paginator.Paginator, error)
	DeleteReaction(ctx context.Context, sc models.Scope, id string) error
}

type CommentRepo interface {
	CreateComment(ctx context.Context, sc models.Scope, opts CreateCommentOptions) (models.Comment, error)
	DetailComment(ctx context.Context, sc models.Scope, id string) (models.Comment, error)
	GetOneComment(ctx context.Context, sc models.Scope, opts GetOneCommentOptions) (models.Comment, error)
	ListComment(ctx context.Context, sc models.Scope, opts ListCommentOptions) ([]models.Comment, error)
	GetComment(ctx context.Context, sc models.Scope, opts GetCommentOptions) ([]models.Comment, paginator.Paginator, error)
	UpdateComment(ctx context.Context, sc models.Scope, opts UpdateCommentOptions) (models.Comment, error)
	DeleteComment(ctx context.Context, sc models.Scope, id string) error
}
