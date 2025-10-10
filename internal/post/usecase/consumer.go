package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post"
)

func (uc impleUsecase) ProcessDeleteCommentMsg(ctx context.Context, sc models.Scope, input post.DeleteCommentMsgInput) error {
	return nil
}

func (uc impleUsecase) ProcessDeleteReactionMsg(ctx context.Context, sc models.Scope, input post.DeleteReactionMsgInput) error {
	return nil
}
