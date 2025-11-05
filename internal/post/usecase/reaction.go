package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/internal/post/repository"
)

func (uc impleUsecase) CreateReaction(ctx context.Context, sc models.Scope, input post.CreateReactionInput) (models.Reaction, error) {
	_, err := uc.repo.Detail(ctx, sc, input.PostID)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.CreateReaction.Detail: %v", err)
		return models.Reaction{}, err
	}

	reaction, err := uc.CreateReaction(ctx, sc, input)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.CreateReaction.CreateReaction: %v", err)
		return models.Reaction{}, err
	}

	return reaction, nil
}

func (uc impleUsecase) DetailReaction(ctx context.Context, sc models.Scope, id string) (models.Reaction, error) {
	reaction, err := uc.repo.DetailReaction(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.DetailReaction.DetailReaction: %v", err)
		return models.Reaction{}, err
	}
	return reaction, nil
}

func (uc impleUsecase) ListReaction(ctx context.Context, sc models.Scope, input post.ListReactionInput) ([]models.Reaction, error) {
	reactions, err := uc.repo.ListReaction(ctx, sc, repository.ListReactionOptions{
		FilterReaction: repository.FilterReaction{
			ID:     input.ID,
			IDs:    input.IDs,
			UserID: input.UserID,
			Type:   input.Type,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.ListReaction.ListReaction: %v", err)
		return []models.Reaction{}, err
	}

	return reactions, nil
}

func (uc impleUsecase) GetReaction(ctx context.Context, sc models.Scope, input post.GetReactionInput) (post.GetReactionOutput, error) {
	reactions, paginator, err := uc.repo.GetReaction(ctx, sc, repository.GetReactionOptions{
		FilterReaction: repository.FilterReaction{
			ID:     input.ID,
			IDs:    input.IDs,
			UserID: input.UserID,
			Type:   input.Type,
		},
		PagQuery: input.PagQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.GetReaction.GetReaction: %v", err)
		return post.GetReactionOutput{}, err
	}
	return post.GetReactionOutput{
		Reactions: reactions,
		Paginator: paginator,
	}, nil
}

func (uc impleUsecase) DeleteReaction(ctx context.Context, sc models.Scope, id string) error {
	err := uc.repo.DeleteReaction(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "reaction.usecase.DeleteReaction.DeleteReaction: %v", err)
		return err
	}
	return nil
}
