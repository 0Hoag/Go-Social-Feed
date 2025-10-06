package usecase

import (
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/internal/post/repository"
	"github.com/hoag/go-social-feed/pkg/log"
)

type impleUsecase struct {
	l    log.Logger
	repo repository.Repository
}

func New(
	l log.Logger,
	repo repository.Repository,
) post.UseCase {
	return &impleUsecase{
		l:    l,
		repo: repo,
	}
}
