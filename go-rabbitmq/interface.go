package rabbitmq

import (
	"github.com/chenyingdi/golang/go-rabbitmq/basic"
	"github.com/streadway/amqp"
	"log"
)

type Modes interface {
	Publish(msg string)
	Consume(handler func(delivery amqp.Delivery))
}

type Mode struct {
	rabbitMQ *basic.RabbitMQ
	Kind     string
}

func NewMode() Modes {
	return &Mode{}
}

func (m *Mode) Publish(msg string) {
	// 1. declare the exchange
	err := m.rabbitMQ.channel.ExchangeDeclare(
		m.rabbitMQ.Exchange,
		m.Kind,
		true,
		false,
		false,
		false,
		nil,
	)
	m.rabbitMQ.failOnError(err, "declare exchange error")

	// 2. send msg
	err = m.rabbitMQ.channel.Publish(
		m.rabbitMQ.Exchange,
		m.rabbitMQ.Key,
		false,
		false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte(msg)},
	)

	m.rabbitMQ.failOnError(err, "publish error")

}

func (m *Mode) Consume(handler func(delivery amqp.Delivery)) {
	// 1. declare the exchange
	err := m.rabbitMQ.channel.ExchangeDeclare(
		m.rabbitMQ.Exchange,
		m.Kind,
		true,
		false,
		false,
		false,
		nil,
	)

	m.rabbitMQ.failOnError(err, "exchange declare error")

	// 2. declare queue
	q, err := m.rabbitMQ.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	m.rabbitMQ.failOnError(err, "queue declare error")

	// 3. bind queque to exchange
	err = m.rabbitMQ.channel.QueueBind(
		q.Name,
		m.rabbitMQ.Key,
		m.rabbitMQ.Exchange,
		false,
		nil,
	)

	m.rabbitMQ.failOnError(err, "binding error")

	// 4. consume msg
	msgs, err := m.rabbitMQ.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// handle msg
			handler(d)
		}
	}()

	log.Printf("[*]Waiting for message...To exit press CTRL+C")
	<-forever
}
