package usecase

import (
	"github.com/hoag/go-social-feed/internal/users"
	"github.com/hoag/go-social-feed/internal/users/repository"
	"github.com/hoag/go-social-feed/pkg/log"
)

type impleUsecase struct {
	l    log.Logger
	repo repository.Repository
}

func New(
	l log.Logger,
	repo repository.Repository,
) users.UseCase {
	return &impleUsecase{
		l:    l,
		repo: repo,
	}
}
