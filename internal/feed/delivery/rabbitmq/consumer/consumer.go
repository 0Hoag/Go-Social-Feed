package consumer

import (
	"log"

	"github.com/hoag/go-social-feed/internal/feed/delivery/rabbitmq"
)

func (c Consumer) Consume() {
	go c.consume(rabbitmq.DeletePostRelationExchange, rabbitmq.DeleteCommentQueueName, c.deleteCommentWorker)
	go c.consume(rabbitmq.DeletePostRelationExchange, rabbitmq.DeleteReactionQueueName, c.deleteReactionWorker)
}

func catchPanic() {
	if r := recover(); r != nil {
		log.Printf("Recovered from panic in goroutine: %v", r)
	}
}
