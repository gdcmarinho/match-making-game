package repository

import (
	"log"

	"github.com/streadway/amqp"
)

func init() {
	startConnection()
	declareQueue()
	openChannel()
}

func Publish() {	
	body := "Someone is trying to play a match"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a requisition")
	log.Printf(" [x] Sent %s\n", body)
}

func startConnection() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
}

func openChannel() {
	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
}

func declareQueue() {
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}