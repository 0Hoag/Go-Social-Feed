package consumer

import (
	"github.com/hoag/go-social-feed/internal/post"
	"github.com/hoag/go-social-feed/pkg/log"
	"github.com/hoag/go-social-feed/pkg/rabbitmq"
)

type Consumer struct {
	l      log.Logger
	postUC post.UseCase
	conn   *rabbitmq.Connection
}

func New(l log.Logger, conn *rabbitmq.Connection) Consumer {
	return Consumer{
		l:    l,
		conn: conn,
	}
}
