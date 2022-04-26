package main

import (
	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672")
	ch, _ := conn.Channel()

	q, _ := ch.QueueDeclare(
		"first.queue", //name string
		true,          //durable bool,
		false,         //autoDelete bool,
		false,         //exclusive bool,
		false,         // noWait bool,
		nil,           //args amqp.Table
	)

	ch.QueueBind(
		q.Name,
		"",
		"amq.fanout",
		false,
		nil,
	)
}
