package producer

import (
	"fmt"

	rabb "github.com/hoag/go-social-feed/internal/feed/delivery/rabbitmq"
	rmqPkg "github.com/hoag/go-social-feed/pkg/rabbitmq"
)

func (p *implProducer) Run() error {
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare(rabb.CreateNotificationExchange)
	if err != nil {
		return err
	}

	p.pushNotiWriter = ch
	return nil
}

// Close closes the producer
func (p *implProducer) Close() {
	p.pushNotiWriter.Close()
	p.deletePostRelationWriter.Close()
}

func (p implProducer) getWriter(exchange rmqPkg.ExchangeArgs) (*rmqPkg.Channel, error) {
	ch, err := p.conn.Channel()
	if err != nil {
		fmt.Println("Error when getting channel")
		return nil, err
	}

	err = ch.ExchangeDeclare(exchange)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (p implProducer) getWriterWithQueue(exchange rmqPkg.ExchangeArgs, queue rmqPkg.QueueArgs) (*rmqPkg.Channel, error) {
	ch, err := p.conn.Channel()
	if err != nil {
		fmt.Println("Error when getting channel")
		return nil, err
	}

	err = ch.ExchangeDeclare(exchange)
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(queue)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(rmqPkg.QueueBindArgs{
		Queue:    queue.Name,
		Exchange: exchange.Name,
	})
	if err != nil {
		return nil, err
	}

	return ch, nil
}
