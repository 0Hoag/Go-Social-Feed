package consumer

import (
	"context"

	"github.com/hoag/go-social-feed/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type WorkerFunc func(msg amqp.Delivery)

func (c Consumer) consume(
	exchange rabbitmq.ExchangeArgs,
	queueName string,
	workerFunc WorkerFunc,
) {
	defer catchPanic()
	ctx := context.Background()

	ch, err := c.conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(exchange)
	if err != nil {
		panic(err)
	}

	q, err := ch.QueueDeclare(rabbitmq.QueueArgs{
		Name:    queueName,
		Durable: true,
	})
	if err != nil {
		panic(err)
	}

	err = ch.QueueBind(rabbitmq.QueueBindArgs{
		Queue:    q.Name,
		Exchange: exchange.Name,
	})
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(rabbitmq.ConsumeArgs{
		Queue: q.Name,
	})
	if err != nil {
		panic(err)
	}

	c.l.Infof(ctx, "Queue %s is being consumed", q.Name)

	var forever chan bool

	go func() {
		for msg := range msgs {
			workerFunc(msg)
		}
	}()

	<-forever
}
