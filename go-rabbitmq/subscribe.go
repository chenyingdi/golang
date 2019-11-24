package rabbitmq

import (
	"github.com/chenyingdi/golang/go-rabbitmq/basic"
	"github.com/streadway/amqp"
)

/*
	Subscribe Mode
*/
type Subscribe struct {
	Mode
}

/*
	create a new subscribe instance
*/
func NewSubscribe(exchange, username, password, host, port, vhost string) *Subscribe {
	var err error
	s := Subscribe{}
	s.Kind = "fanout"
	s.rabbitMQ = basic.NewRabbitMQ("", exchange, "")

	s.rabbitMQ.InitUrl(username, password, host, port, vhost)
	s.rabbitMQ.Conn, err = amqp.Dial(s.rabbitMQ.Url.ParseUrl())

	s.rabbitMQ.FailOnError(err, "dial error")

	s.rabbitMQ.Channel, err = s.rabbitMQ.Conn.Channel()

	s.rabbitMQ.FailOnError(err, "get channel error")

	return &s
}


