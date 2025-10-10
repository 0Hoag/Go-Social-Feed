package producer

import (
	"context"
	"encoding/json"

	rabb "github.com/hoag/go-social-feed/internal/feed/delivery/rabbitmq"
	"github.com/hoag/go-social-feed/pkg/rabbitmq"
)

func (p implProducer) PublishNotiMsg(ctx context.Context, msg rabb.PublishNotiMsg) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return p.pushNotiWriter.Publish(ctx, rabbitmq.PublishArgs{
		Exchange: rabb.CreateNotificationExchange.Name,
		Msg: rabbitmq.Publishing{
			Body:        body,
			ContentType: rabbitmq.ContentTypePlainText,
		},
	})
}

func (p implProducer) PublishDeletePostRelationMsg(ctx context.Context, msg rabb.DeletePostRelationMsg) error {
	body, err := json.Marshal(msg)
	if err != nil {
		p.l.Errorf(ctx, "feed.delivery.rabbitmq.producer.PublishDeletePostRelationMsg.Marshal: %v", err)
		return err
	}

	return p.deletePostRelationWriter.Publish(ctx, rabbitmq.PublishArgs{
		Exchange: rabb.DeletePostRelationExchange.Name,
		Msg: rabbitmq.Publishing{
			Body:        body,
			ContentType: rabbitmq.ContentTypePlainText,
		},
	})
}
