package rabbitmq

import (
	"github.com/chenyingdi/golang/go-rabbitmq/basic"
	"github.com/streadway/amqp"
)

type Routing struct {
	Mode
}

func NewRouting(exchange, routingKey, username, password, host, port, vhost string) *Routing {
	r := Routing{}
	r.Kind = "direct"
	r.rabbitMQ = basic.NewRabbitMQ("", exchange, routingKey)

	var err error

	r.rabbitMQ.InitUrl(username, password, host, port, vhost)
	r.rabbitMQ.Conn, err = amqp.Dial(r.rabbitMQ.Url.ParseUrl())

	r.rabbitMQ.FailOnError(err, "dialing error")

	r.rabbitMQ.Channel, err = r.rabbitMQ.Conn.Channel()

	r.rabbitMQ.FailOnError(err, "get channel error")

	return &r
}
