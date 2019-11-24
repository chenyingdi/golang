package rabbitmq

import (
	"github.com/chenyingdi/golang/go-rabbitmq/basic"
	"github.com/streadway/amqp"
)

type Topic struct {
	Mode
}

func NewTopic(exchange, routingKey, username, password, host, port, vhost string) *Topic {
	var err error
	t := Topic{}
	t.Kind = "topic"
	t.rabbitMQ = basic.NewRabbitMQ("", exchange, routingKey)

	t.rabbitMQ.InitUrl(username, password, host, port, vhost)
	t.rabbitMQ.Conn, err = amqp.Dial(t.rabbitMQ.Url.ParseUrl())

	t.rabbitMQ.FailOnError(err, "dialing error")

	t.rabbitMQ.Channel, err = t.rabbitMQ.Conn.Channel()

	t.rabbitMQ.FailOnError(err, "get channel error")

	return &t
}
