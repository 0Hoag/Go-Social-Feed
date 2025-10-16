package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/internal/post/repository"
)

func (uc impleUsecase) Create(ctx context.Context, sc models.Scope, input post.CreateInput) (models.Post, error) {
	post, err := uc.repo.Create(ctx, sc, repository.CreateOptions{
		Pin:          input.Pin,
		Content:      input.Content,
		FileIDs:      input.FileIDs,
		TaggedTarget: input.TaggedTarget,
		Permission:   input.Permission,
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.Create.Create: %v", err)
		return models.Post{}, err
	}

	if len(post.TaggedTarget) > 0 {
		err = uc.handleCreatePostNotification(ctx, sc, post)
		if err != nil {
			uc.l.Errorf(ctx, "post.usecase.Create.handleCreatePostNotification : %v", err)
			return models.Post{}, nil
		}
	}

	return post, nil
}

func (uc impleUsecase) Detail(ctx context.Context, sc models.Scope, id string) (models.Post, error) {
	post, err := uc.repo.Detail(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.Detail.Detail: %v", err)
		return models.Post{}, err
	}
	return post, nil
}

func (uc impleUsecase) List(ctx context.Context, sc models.Scope, input post.ListInput) ([]models.Post, error) {
	posts, err := uc.repo.List(ctx, sc, repository.ListOptions{
		Filter: repository.Filter{
			ID:       input.ID,
			IDs:      input.IDs,
			Pin:      input.Pin,
			AuthorID: input.AuthorID,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.List.List: %v", err)
		return []models.Post{}, err
	}
	return posts, nil
}

func (uc impleUsecase) Get(ctx context.Context, sc models.Scope, input post.GetInput) (post.GetOutput, error) {
	posts, paginator, err := uc.repo.Get(ctx, sc, repository.GetOptions{
		Filter: repository.Filter{
			ID:       input.ID,
			IDs:      input.IDs,
			Pin:      input.Pin,
			AuthorID: input.AuthorID,
		},
		PagQuery: input.PagQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.Get.Get: %v", err)
		return post.GetOutput{}, err
	}
	return post.GetOutput{
		Posts:     posts,
		Paginator: paginator,
	}, nil
}

func (uc impleUsecase) Update(ctx context.Context, sc models.Scope, input post.UpdateInput) error {
	post, err := uc.repo.Detail(ctx, sc, input.ID)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.Update.Detail: %v", err)
		return err
	}

	err = uc.repo.Update(ctx, sc, repository.UpdateOptions{
		Post:         post,
		Content:      input.Content,
		FileIDs:      input.FileIDs,
		Pin:          input.Pin,
		TaggedTarget: input.TaggedTarget,
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.Update.Update: %v", err)
		return err
	}
	return nil
}

func (uc impleUsecase) Delete(ctx context.Context, sc models.Scope, id string) error {
	err := uc.repo.Delete(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.Delete.Delete: %v", err)
		return err
	}
	return nil
}
