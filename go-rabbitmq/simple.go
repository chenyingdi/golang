package rabbitmq

import (
	"github.com/chenyingdi/golang/go-rabbitmq/basic"
	"github.com/streadway/amqp"
	"log"
)

type Simple struct {
	Mode
}

/*
	create a new rabbitMQ instance for simple mode or work mode
*/
func NewSimple(queueName, username, password, host, port, vhost string) Modes {
	var err error
	b := Simple{}
	b.rabbitMQ = basic.NewRabbitMQ(queueName, "", "")
	b.rabbitMQ.InitUrl(username, password, host, port, vhost)
	b.rabbitMQ.Conn, err = amqp.Dial(b.rabbitMQ.Url.ParseUrl())
	b.rabbitMQ.FailOnError(err, "dial error")

	b.rabbitMQ.Channel, err = b.rabbitMQ.Conn.Channel()
	b.rabbitMQ.FailOnError(err, "get channel error")

	return &b
}

/*
	simple publish
*/
func (s *Simple) Publish(msg string) {
	// 1. declare the queue
	_, err := s.rabbitMQ.Channel.QueueDeclare(
		s.rabbitMQ.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	s.rabbitMQ.FailOnError(err, "queue declare error")

	// 2. send msg
	s.rabbitMQ.Channel.Publish(
		s.rabbitMQ.Exchange,
		s.rabbitMQ.QueueName,
		false,
		false,
		amqp.Publishing{ContentType:"text/plain", Body:[]byte(msg)},
	)
}

/*
	simple consume
 */
func (s *Simple) Consume(handler func(delivery amqp.Delivery)) {
	// 1. declare the queue
	_, err := s.rabbitMQ.Channel.QueueDeclare(
		s.rabbitMQ.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	s.rabbitMQ.FailOnError(err, "queue declare error")

	// 2. receive msg
	msg, err := s.rabbitMQ.Channel.Consume(
		s.rabbitMQ.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	s.rabbitMQ.FailOnError(err, "consume error")

	forever := make(chan bool)

	go func() {
		for d := range msg {
			// msg handler
			handler(d)
		}
	}()

	log.Printf("[*]Waiting for message...")
	<- forever
}
