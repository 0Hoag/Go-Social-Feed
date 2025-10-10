package usecase

import (
	"github.com/hoag/go-social-feed/internal/feed/delivery/rabbitmq/producer"
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/internal/post/repository"
	"github.com/hoag/go-social-feed/pkg/log"
)

type impleUsecase struct {
	l    log.Logger
	prod producer.Producer
	repo repository.Repository
}

func New(
	l log.Logger,
	prod producer.Producer,
	repo repository.Repository,
) post.UseCase {
	return &impleUsecase{
		l:    l,
		prod: prod,
		repo: repo,
	}
}
