package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/follow"
	"github.com/hoag/go-social-feed/internal/follow/repository"
	"github.com/hoag/go-social-feed/internal/models"
)

func (uc impleUsecase) Create(ctx context.Context, sc models.Scope, input follow.CreateInput) (models.Follow, error) {
	_, err := uc.userUC.Detail(ctx, sc, input.FolloweeID)
	if err != nil {
		uc.l.Errorf(ctx, "follow.usecase.Create.Detail: %v", err)
		return models.Follow{}, err
	}

	follow, err := uc.Create(ctx, sc, input)
	if err != nil {
		uc.l.Errorf(ctx, "follow.usecase.Create.Create: %v", err)
		return models.Follow{}, err
	}

	return follow, nil
}

func (uc impleUsecase) Detail(ctx context.Context, sc models.Scope, id string) (models.Follow, error) {
	follow, err := uc.repo.Detail(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "follow.usecase.Detail.Detail: %v", err)
		return models.Follow{}, err
	}
	return follow, nil
}

func (uc impleUsecase) List(ctx context.Context, sc models.Scope, input follow.ListInput) ([]models.Follow, error) {
	follows, err := uc.repo.List(ctx, sc, repository.ListOptions{
		Filter: repository.Filter{
			ID:         input.ID,
			IDs:        input.IDs,
			AuthorID:   input.AuthorID,
			FolloweeID: input.FolloweeID,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "follow.usecase.List.List: %v", err)
		return []models.Follow{}, err
	}
	return follows, nil
}

func (uc impleUsecase) Get(ctx context.Context, sc models.Scope, input follow.GetInput) (follow.GetOutput, error) {
	follows, paginator, err := uc.repo.Get(ctx, sc, repository.GetOptions{
		Filter: repository.Filter{
			ID:         input.ID,
			IDs:        input.IDs,
			AuthorID:   input.AuthorID,
			FolloweeID: input.FolloweeID,
		},
		PagQuery: input.PagQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "follow.usecase.Get.Get: %v", err)
		return follow.GetOutput{}, err
	}
	return follow.GetOutput{
		Follows:   follows,
		Paginator: paginator,
	}, nil
}

func (uc impleUsecase) Delete(ctx context.Context, sc models.Scope, id string) error {
	err := uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "follow.usecase.Delete.Delete: %v", err)
		return err
	}
	return nil
}
