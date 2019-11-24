package basic

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)


type RabbitMQ struct {
	Conn      *amqp.Connection
	Channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	Url       *Url
}

/*
	create a new rabbitMQ instance
 */
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ{
	return &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
	}
}

/*
	init url
 */
func (r *RabbitMQ) InitUrl(username, password, host, port, vhost string)  {
	r.Url = NewUrl(username, password, host, port, vhost)
}

/*
	destroy
 */
func (r *RabbitMQ) Destroy()  {
	r.Channel.Close()
	r.Conn.Close()
}


/*
	error handler
 */
func (r *RabbitMQ) FailOnError(err error, msg string)  {
	if err != nil{
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

/*
	publish
 */
func (r *RabbitMQ) Publish()  {
	
}