package main

import (
	"bytes"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest@localhost:5672/")
	failOnErrors(err, "Failed to connect Rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnErrors(err, "Failed to open a Channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello cak", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusif
		false,       // no-wait
		nil,         // arguments
	)
	failOnErrors(err, "Failed to Declare a queue")

	msgs, err := ch.Consume(
		q.Name, // quee
		"",     // consumer
		true,   // auto-act
		false,  // exclusif
		false,  // no-local
		false,  // no-wait
		nil,
	)
	failOnErrors(err, "Failed to register a consumer")

	foever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a massage: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Println("Done")
			d.Ack(false)
		}
	}()

	log.Println(" [*] Waiting For Massages. To Exit press CTRL+C")

	<-foever
}

func failOnErrors(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
