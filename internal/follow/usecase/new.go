package usecase

import (
	"github.com/hoag/go-social-feed/internal/follow"
	"github.com/hoag/go-social-feed/internal/follow/repository"
	"github.com/hoag/go-social-feed/internal/users"
	"github.com/hoag/go-social-feed/pkg/log"
)

type impleUsecase struct {
	l      log.Logger
	userUC users.UseCase
	repo   repository.Repository
}

func New(
	l log.Logger,
	userUC users.UseCase,
	repo repository.Repository,
) follow.UseCase {
	return &impleUsecase{
		l:      l,
		userUC: userUC,
		repo:   repo,
	}
}
