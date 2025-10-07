package usecase

import (
	"github.com/hoag/go-social-feed/config"
	"github.com/hoag/go-social-feed/internal/auth"
	"github.com/hoag/go-social-feed/internal/users"
	"github.com/hoag/go-social-feed/pkg/log"
)

type impleUsecase struct {
	l      log.Logger
	cfg    *config.Config
	userUC users.UseCase
}

func New(
	l log.Logger,
	cfg *config.Config,
	userUC users.UseCase,
) auth.UseCase {
	return &impleUsecase{
		l:      l,
		cfg:    cfg,
		userUC: userUC,
	}
}
