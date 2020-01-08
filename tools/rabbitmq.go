package tools

import (
	"os"

	"github.com/streadway/amqp"
)

const (
	EXNAME = "messages"
	QNAME  = "saveMessages"
	KEY    = "message.created"
)

type RabbitClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitClient() *RabbitClient {
	conn := initConn()
	chann := initChannel(conn)
	return &RabbitClient{
		Conn:    conn,
		Channel: chann,
	}
}

func initConn() *amqp.Connection {
	var amqpUrl string
	if url := os.Getenv("AMQP_URL"); url == "" {
		amqpUrl = "amqp://chatapp:chatpass@rabbitmq:5672"
	} else {
		amqpUrl = url
	}

	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		panic(err)
	}

	return conn
}

func initChannel(conn *amqp.Connection) *amqp.Channel {
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return channel
}

func (c *RabbitClient) SetUp() {
	err := c.Channel.ExchangeDeclare(EXNAME, amqp.ExchangeDirect, true, false, false, false, nil)
	handleErr(err)

	_, err = c.Channel.QueueDeclare(QNAME, true, false, false, false, nil)
	handleErr(err)

	err = c.Channel.QueueBind(QNAME, KEY, EXNAME, false, nil)
	handleErr(err)
}

func (c *RabbitClient) Publish(exchange, key string, message []byte) {
	msg := amqp.Publishing{
		Body: message,
	}
	err := c.Channel.Publish(exchange, key, false, false, msg)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
