package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Cannot connect in Rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed Open Channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello cak", // name
		false,       // durable
		false,       // deleted when unused
		false,       // ekslusif
		false,       // no-wait
		nil,         // argument
	)
	failOnError(err, "Failed to declare a queue")
	for i := 0; i < 1000; i++ {
		body := bodyFrom(os.Args, i)

		err = ch.Publish(
			"",     // exchange
			q.Name, //routing key
			false,  //mandatory
			false,  // immediati
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "tect/plain",
				Body:         []byte(body),
			})
	}

	failOnError(err, "Failed to publish Massage")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

func bodyFrom(args []string, i int) string {
	var s string

	if (len(args) < 2) || os.Args[1] == "" {
		s = "Hello sempak ijo lek" + strconv.Itoa(i)
	} else {
		s = strings.Join(args[1:], " ")
	}

	return s
}
