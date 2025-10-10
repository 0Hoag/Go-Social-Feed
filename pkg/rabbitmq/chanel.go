package rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Channel struct {
	conn       *Connection
	ch         *amqp.Channel
	reconnects []chan bool
}

func (c Channel) ExchangeDeclare(exc ExchangeArgs) error {
	return c.ch.ExchangeDeclare(exc.spread())
}

func (c Channel) QueueDeclare(queue QueueArgs) (amqp.Queue, error) {
	return c.ch.QueueDeclare(queue.spread())
}

func (c Channel) QueueBind(queueBind QueueBindArgs) error {
	return c.ch.QueueBind(queueBind.spread())
}

func (c Channel) Publish(ctx context.Context, publish PublishArgs) error {
	return c.ch.PublishWithContext(publish.spread(ctx))
}

func (c Channel) Consume(consume ConsumeArgs) (<-chan amqp.Delivery, error) {
	return c.ch.Consume(consume.spread())
}

func (c Channel) Close() error {
	return c.ch.Close()
}

func (c *Channel) listenNotifyReconnect() {
	reconnNoti := make(chan bool)
	c.conn.notifyReconect(reconnNoti)

	go func() {
		for {
			<-reconnNoti

			log.Println("Retry creating RabbitMQ channel...")
			channel, err := c.conn.channel()
			if err != nil {
				log.Printf("RabbitMQ channel failed: %v\n", err)
				continue
			}
			c.Close()
			c.ch = channel
		}
	}()
}

func (c *Channel) NotifyReconect(reciver chan bool) <-chan bool {
	c.reconnects = append(c.reconnects, reciver)
	return reciver
}
