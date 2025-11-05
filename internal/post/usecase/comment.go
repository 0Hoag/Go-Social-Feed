package usecase

import (
	"context"

	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/internal/post/repository"
)

func (uc impleUsecase) CreateComment(ctx context.Context, sc models.Scope, input post.CreateCommentInput) (models.Comment, error) {
	_, err := uc.repo.DetailComment(ctx, sc, input.PostID)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.CreateComment.DetailComment: %v", err)
		return models.Comment{}, err
	}

	comment, err := uc.repo.CreateComment(ctx, sc, repository.CreateCommentOptions{
		PostID:  input.PostID,
		Content: input.Content,
		Attach:  input.Attach,
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.CreateComment.CreateComment: %v", err)
		return models.Comment{}, err
	}

	return comment, nil
}

func (uc impleUsecase) DetailComment(ctx context.Context, sc models.Scope, id string) (models.Comment, error) {
	comment, err := uc.repo.DetailComment(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.DetailComment.DetailComment: %v", err)
		return models.Comment{}, err
	}
	return comment, nil
}

func (uc impleUsecase) ListComment(ctx context.Context, sc models.Scope, input post.ListCommentInput) ([]models.Comment, error) {
	comments, err := uc.repo.ListComment(ctx, sc, repository.ListCommentOptions{
		FilterComment: repository.FilterComment{
			ID:  input.ID,
			IDs: input.IDs,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.ListComment.ListComment: %v", err)
		return []models.Comment{}, err
	}
	return comments, nil
}

func (uc impleUsecase) GetComment(ctx context.Context, sc models.Scope, input post.GetCommentInput) (post.GetCommentOutput, error) {
	comments, paginator, err := uc.repo.GetComment(ctx, sc, repository.GetCommentOptions{
		FilterComment: repository.FilterComment{
			ID:  input.ID,
			IDs: input.IDs,
		},
		PagQuery: input.PagQuery,
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.GetComment.GetComment: %v", err)
		return post.GetCommentOutput{}, err
	}
	return post.GetCommentOutput{
		Comments:  comments,
		Paginator: paginator,
	}, nil
}

func (uc impleUsecase) UpdateComment(ctx context.Context, sc models.Scope, input post.UpdateCommentInput) (models.Comment, error) {
	m, err := uc.repo.GetOneComment(ctx, sc, repository.GetOneCommentOptions{
		FilterComment: repository.FilterComment{
			PostID: input.PostID,
		},
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.UpdateComment.GetOneComment: %v", err)
		return models.Comment{}, err
	}

	m, err = uc.repo.UpdateComment(ctx, sc, repository.UpdateCommentOptions{
		Comment: m,
		Content: input.Content,
		Attach:  input.Attach,
	})
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.UpdateComment.UpdateComment: %v", err)
		return models.Comment{}, err
	}

	return m, nil
}

func (uc impleUsecase) DeleteComment(ctx context.Context, sc models.Scope, id string) error {
	err := uc.repo.DeleteComment(ctx, sc, id)
	if err != nil {
		uc.l.Errorf(ctx, "post.usecase.DeleteComment.DeleteComment: %v", err)
		return err
	}
	return nil
}
