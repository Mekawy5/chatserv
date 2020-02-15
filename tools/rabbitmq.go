package tools

import (
	"os"

	"github.com/streadway/amqp"
)

const (
	MSGEXC = "messages"
	MSGQ   = "saveMessages"
	MSGKEY = "message.created"
	CHTEXC = "chats"
	CHTQ   = "saveChats"
	CHTKEY = "chats.created"
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

func (c *RabbitClient) SetUpMsg() {
	// messages queue setup
	err := c.Channel.ExchangeDeclare(MSGEXC, amqp.ExchangeDirect, true, false, false, false, nil)
	handleErr(err)
	_, err = c.Channel.QueueDeclare(MSGQ, true, false, false, false, nil)
	handleErr(err)
	err = c.Channel.QueueBind(MSGQ, MSGKEY, MSGEXC, false, nil)
	handleErr(err)
}

func (c *RabbitClient) SetUpCht() {
	// chats queue setup
	err := c.Channel.ExchangeDeclare(CHTEXC, amqp.ExchangeDirect, true, false, false, false, nil)
	handleErr(err)
	_, err = c.Channel.QueueDeclare(CHTQ, true, false, false, false, nil)
	handleErr(err)
	err = c.Channel.QueueBind(CHTQ, CHTKEY, CHTEXC, false, nil)
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
