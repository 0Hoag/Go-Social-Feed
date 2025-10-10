package rabbitmq

import "github.com/hoag/go-social-feed/pkg/rabbitmq"

const (
	DeletePostRelationExchangeName = "connect_delete_post_relation_exc"

	CreateNotificationExchangeName = "notification_create_exc"
	CreateNotificationQueueName    = "notification_create"

	DeleteReactionQueueName = "connect_delete_reaction"
	DeleteCommentQueueName  = "connect_delete_comment"
)

var (
	DeletePostRelationExchange = rabbitmq.ExchangeArgs{
		Name:       DeletePostRelationExchangeName,
		Type:       rabbitmq.ExchangeTypeFanout,
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
	}

	CreateNotificationExchange = rabbitmq.ExchangeArgs{
		Name:       CreateNotificationExchangeName,
		Type:       rabbitmq.ExchangeTypeFanout,
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
	}
)
