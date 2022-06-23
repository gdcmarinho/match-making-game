package main

import (
	"log"
    "net/http"

    "github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type player struct {
    ID     		string  `json:"id"`
    Nickname  	string  `json:"nickname"`
    Tag 		string  `json:"tag"`
    Rank  		int 	`json:"rank"`
}

var players = []player{
    {ID: "1", Nickname: "Marinho", Tag: "GMX", Rank: 1},
    {ID: "1", Nickname: "4Queijos", Tag: "BRL", Rank: 2},
    {ID: "1", Nickname: "Disco-Lee", Tag: "NFT", Rank: 3},
}

func findMatch(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, players)
}

func main() {
    router := gin.Default()
    router.GET("/players", findMatch)

    router.Run("localhost:8080")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

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

func failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}