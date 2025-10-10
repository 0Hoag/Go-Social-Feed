package rabbitmq

import (
	"errors"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	retryConnectionDelay   = 2 * time.Second
	retryConnectionTimeout = 20 * time.Second
)

var (
	ErrConnectionTimeout = errors.New("connection timeout")
)

type Connection struct {
	url                 string
	retryWithoutTimeout bool
	conn                *amqp.Connection
	isRetrying          bool
	reconnects          []chan bool
}

func Dial(url string, retryWithoutTimeout bool) (Connection, error) {
	conn := Connection{
		url:                 url,
		retryWithoutTimeout: retryWithoutTimeout,
	}

	if err := conn.connect(); err != nil {
		return Connection{}, err
	}

	return conn, nil
}

func (c *Connection) Close() {
	if c.conn != nil {
		c.conn.Close()
	}

	c.isRetrying = false
}

func (c Connection) IsReady() bool {
	return c.conn != nil && !c.conn.IsClosed()
}

func (c Connection) IsClosed() bool {
	return !c.IsReady() && !c.isRetrying
}

func (c Connection) dial(url string, connChan chan *amqp.Connection, cancelChan chan bool) {
	count := 0
	for {
		select {
		case <-cancelChan:
			return
		default:
			log.Printf("Trying to connect to RabbitMQ, attempt: %d ...\n", count+1)
			conn, err := amqp.Dial(url)

			if err != nil {
				log.Printf("--- RabbitMQ connection failed: %v\n", err)
				time.Sleep(retryConnectionDelay)
				count++
				continue
			}

			connChan <- conn
			return
		}
	}
}

func (c *Connection) connectWithoutTimeout() error {
	connChan := make(chan *amqp.Connection)

	go c.dial(c.url, connChan, make(chan bool))

	conn := <-connChan
	log.Println("RabbitMQ connection established")
	c.conn = conn
	c.listenNotifyClose()
	return nil
}

func (c *Connection) connect() error {
	connChan := make(chan *amqp.Connection)
	cancelChan := make(chan bool)

	go c.dial(c.url, connChan, cancelChan)

	select {
	case conn := <-connChan:
		log.Println("RabbitMQ connection established")
		c.conn = conn
		c.listenNotifyClose()
		return nil
	case <-time.After(retryConnectionTimeout):
		cancelChan <- true
		return ErrConnectionTimeout
	}
}

func (c *Connection) listenNotifyClose() {
	fn := c.connect
	if c.retryWithoutTimeout {
		fn = c.connectWithoutTimeout
	}

	notifyClose := make(chan *amqp.Error)
	c.conn.NotifyClose(notifyClose)

	go func() {
		for {
			err := <-notifyClose
			if err != nil {
				c.conn = nil
				c.isRetrying = true
				log.Printf("RabbitMQ connection closed: %v\n", err)

				if err := fn(); err != nil {
					log.Printf("RabbitMQ connection failed: %v\n", err)
				}

				for _, reconnect := range c.reconnects {
					reconnect <- true
				}

				c.isRetrying = false
				return
			}
		}
	}()
}

func (c Connection) channel() (*amqp.Channel, error) {
	ch, err := c.conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (c *Connection) Channel() (*Channel, error) {
	ch, err := c.channel()
	if err != nil {
		return nil, err
	}

	channel := Channel{
		conn: c,
		ch:   ch,
	}

	channel.listenNotifyReconnect()

	return &channel, nil
}

func (c *Connection) notifyReconect(reciver chan bool) <-chan bool {
	c.reconnects = append(c.reconnects, reciver)
	return reciver
}
