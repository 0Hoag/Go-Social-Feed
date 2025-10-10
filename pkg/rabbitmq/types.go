package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ExchangeArgs is a struct that contains the arguments for creating an exchange
type ExchangeArgs struct {
	Name       string
	Type       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       map[string]interface{}
}

func (e ExchangeArgs) spread() (name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) {
	return e.Name, e.Type, e.Durable, e.AutoDelete, e.Internal, e.NoWait, e.Args
}

// QueueArgs is a struct that contains the arguments for creating a queue
type QueueArgs struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       map[string]interface{}
}

func (q QueueArgs) spread() (name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) {
	return q.Name, q.Durable, q.AutoDelete, q.Exclusive, q.NoWait, q.Args
}

// Publishing is a struct that contains the arguments for publishing a message
type Publishing = amqp.Publishing

// PublishArgs is a struct that contains the arguments for publishing a message
type PublishArgs struct {
	Exchange   string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Msg        Publishing
}

func (p PublishArgs) spread(ctx context.Context) (c context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) {
	return ctx, p.Exchange, p.RoutingKey, p.Mandatory, p.Immediate, p.Msg
}

// ConsumeArgs is a struct that contains the arguments for consuming a message
type ConsumeArgs struct {
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      map[string]interface{}
}

func (c ConsumeArgs) spread() (queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) {
	return c.Queue, c.Consumer, c.AutoAck, c.Exclusive, c.NoLocal, c.NoWait, c.Args
}

// QueueBindArgs is a struct that contains the arguments for binding a queue
type QueueBindArgs struct {
	Queue      string
	Exchange   string
	RoutingKey string
	NoWait     bool
	Args       map[string]interface{}
}

func (q QueueBindArgs) spread() (queue, key, exchange string, noWait bool, args amqp.Table) {
	return q.Queue, q.RoutingKey, q.Exchange, q.NoWait, q.Args
}
