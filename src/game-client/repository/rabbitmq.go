package repository

import (
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var q amqp.Queue

func init() {
	conn = startConnection()
	ch = openChannel(conn)
	q = declareQueue(ch)
}

func Publish() {	
	body := "Someone is trying to play a match"
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a requisition", err)
	}

	log.Printf(" [x] Sent %s\n", body)
}

func startConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	defer conn.Close()
	return conn
}

func openChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	defer ch.Close()
	return ch
}

func declareQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"hello", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	return q
}