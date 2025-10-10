package consumer

import (
	"context"
	"encoding/json"

	"github.com/hoag/go-social-feed/internal/feed/delivery/rabbitmq"
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c Consumer) deleteCommentWorker(d amqp.Delivery) {
	ctx := context.Background()

	var msg rabbitmq.DeletePostRelationMsg
	err := json.Unmarshal(d.Body, &msg)
	if err != nil {
		c.l.Warnf(ctx, "feed.delivery.rabbitmq.consumer.Unmarshal: %v", err)
		d.Nack(false, true)
		return
	}

	sc := models.Scope{
		UserID: msg.UserID,
	}

	err = c.postUC.ProcessDeleteCommentMsg(ctx, sc, post.DeleteCommentMsgInput{
		PostID: msg.PostID,
	})
	if err != nil {
		c.l.Errorf(ctx, "feed.delivery.rabbitmq.consumer.DeletePostRelation: %v", err)
		d.Nack(false, true)
		return
	}

	d.Ack(false)
}

func (c Consumer) deleteReactionWorker(d amqp.Delivery) {
	ctx := context.Background()

	var msg rabbitmq.DeletePostRelationMsg
	err := json.Unmarshal(d.Body, &msg)
	if err != nil {
		c.l.Warnf(ctx, "feed.delivery.rabbitmq.consumer.Unmarshal: %v", err)
		d.Nack(false, true)
		return
	}

	sc := models.Scope{
		UserID: msg.UserID,
	}

	err = c.postUC.ProcessDeleteReactionMsg(ctx, sc, post.DeleteReactionMsgInput{
		PostID: msg.PostID,
	})
	if err != nil {
		c.l.Errorf(ctx, "feed.delivery.rabbitmq.consumer.DeletePostRelation: %v", err)
		d.Nack(false, true)
		return
	}

	d.Ack(false)
}
