package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Failed to connect Rabbitmq : ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a Channel : ", err)
	}

	q, err := ch.QueueDeclare(
		"TRY-MESSAGING", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusif
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Println("Failed to Declare a queue : ", err)
	}

	msgs, err := ch.Consume(
		q.Name, // quee
		"",     // consumer
		false,  // auto-act
		false,  // exclusif
		false,  // no-local
		false,  // no-wait
		nil,
	)
	if err != nil {
		log.Println("Failed to register a consumer : ", err)
	}

	foever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a massage: %s", d.Body)
			d.Ack(false)
		}
	}()

	log.Println(" [*] Waiting For Massages. To Exit press CTRL+C")

	<-foever
}
