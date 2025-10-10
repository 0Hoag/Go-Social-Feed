package usecase

import (
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/internal/reaction"
	"github.com/hoag/go-social-feed/internal/reaction/repository"
	"github.com/hoag/go-social-feed/pkg/log"
)

type impleUsecase struct {
	l      log.Logger
	postUC post.UseCase
	repo   repository.Repository
}

func New(
	l log.Logger,
	postUC post.UseCase,
	repo repository.Repository,
) reaction.UseCase {
	return &impleUsecase{
		l:      l,
		postUC: postUC,
		repo:   repo,
	}
}
