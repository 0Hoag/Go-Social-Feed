package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/reaction"
	"github.com/hoag/go-social-feed/internal/reaction/repository"
)

func (uc impleUsecase) Create(ctx context.Context, sc models.Scope, input reaction.CreateInput) (models.Reaction, error) {
	_, err := uc.postUC.Detail(ctx, sc, input.PostID)
	if err != nil {
		uc.l.Errorf(ctx, "reaction.usecase.Create.Detail: %v", err)
		return models.Reaction{}, err
	}

	reaction, err := uc.Create(ctx, sc, input)
	if err != nil {
		uc.l.Errorf(ctx, "reaction.usecase.Create.Create: %v", err)
		return models.Reaction{}, err
	}

	return reaction, nil
}

func (uc impleUsecase) Detail(ctx context.Context, sc models.Scope, id string) (models.Reaction, error) {
	reaction, err := uc.repo.Detail(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "reaction.usecase.Detail.Detail: %v", err)
		return models.Reaction{}, err
	}
	return reaction, nil
}

func (uc impleUsecase) List(ctx context.Context, sc models.Scope, input reaction.ListInput) ([]models.Reaction, error) {
	reactions, err := uc.repo.List(ctx, sc, repository.ListOptions{
		Filter: repository.Filter{
			ID:     input.ID,
			IDs:    input.IDs,
			UserID: input.UserID,
			Type:   input.Type,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "reaction.usecase.List.List: %v", err)
		return []models.Reaction{}, err
	}
	return reactions, nil
}

func (uc impleUsecase) Get(ctx context.Context, sc models.Scope, input reaction.GetInput) (reaction.GetOutput, error) {
	reactions, paginator, err := uc.repo.Get(ctx, sc, repository.GetOptions{
		Filter: repository.Filter{
			ID:     input.ID,
			IDs:    input.IDs,
			UserID: input.UserID,
			Type:   input.Type,
		},
		PagQuery: input.PagQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "reaction.usecase.Get.Get: %v", err)
		return reaction.GetOutput{}, err
	}
	return reaction.GetOutput{
		Reactions: reactions,
		Paginator: paginator,
	}, nil
}

func (uc impleUsecase) Delete(ctx context.Context, sc models.Scope, id string) error {
	err := uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "reaction.usecase.Delete.Delete: %v", err)
		return err
	}
	return nil
}
