package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/comment"
	"github.com/hoag/go-social-feed/internal/comment/repository"
	"github.com/hoag/go-social-feed/internal/models"
)

func (uc impleUsecase) Create(ctx context.Context, sc models.Scope, input comment.CreateInput) (models.Comment, error) {
	_, err := uc.postUC.Detail(ctx, sc, input.PostID)
	if err != nil {
		uc.l.Errorf(ctx, "comment.usecase.Create.Detail: %v", err)
		return models.Comment{}, err
	}

	comment, err := uc.repo.Create(ctx, sc, repository.CreateOptions{
		PostID:  input.PostID,
		Content: input.Content,
		Attach:  input.Attach,
	})
	if err != nil {
		uc.l.Errorf(ctx, "comment.usecase.Create.Create: %v", err)
		return models.Comment{}, err
	}

	return comment, nil
}

func (uc impleUsecase) Detail(ctx context.Context, sc models.Scope, id string) (models.Comment, error) {
	comment, err := uc.repo.Detail(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "comment.usecase.Detail.Detail: %v", err)
		return models.Comment{}, err
	}
	return comment, nil
}

func (uc impleUsecase) List(ctx context.Context, sc models.Scope, input comment.ListInput) ([]models.Comment, error) {
	comments, err := uc.repo.List(ctx, sc, repository.ListOptions{
		Filter: repository.Filter{
			ID:  input.ID,
			IDs: input.IDs,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "comment.usecase.List.List: %v", err)
		return []models.Comment{}, err
	}
	return comments, nil
}

func (uc impleUsecase) Get(ctx context.Context, sc models.Scope, input comment.GetInput) (comment.GetOutput, error) {
	comments, paginator, err := uc.repo.Get(ctx, sc, repository.GetOptions{
		Filter: repository.Filter{
			ID:  input.ID,
			IDs: input.IDs,
		},
		PagQuery: input.PagQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "comment.usecase.Get.Get: %v", err)
		return comment.GetOutput{}, err
	}
	return comment.GetOutput{
		Comments:  comments,
		Paginator: paginator,
	}, nil
}

func (uc impleUsecase) Update(ctx context.Context, sc models.Scope, input comment.UpdateInput) (models.Comment, error) {
	m, err := uc.repo.GetOne(ctx, sc, repository.GetOneOptions{
		Filter: repository.Filter{
			PostID: input.PostID,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "comment.usecase.Update.Detail: %v", err)
		return models.Comment{}, err
	}

	m, err = uc.repo.Update(ctx, sc, repository.UpdateOptions{
		Comment: m,
		Content: input.Content,
		Attach:  input.Attach,
	})
	if err != nil {
		uc.l.Errorf(ctx, "comment.usecase.Update.Detail: %v", err)
		return models.Comment{}, err
	}

	return m, nil
}

func (uc impleUsecase) Delete(ctx context.Context, sc models.Scope, id string) error {
	err := uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "comment.usecase.Delete.Delete: %v", err)
		return err
	}
	return nil
}
