package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Cannot connect in Rabbitmq : ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed Open Channel : ", err)
	}

	q, err := ch.QueueDeclare(
		"COBALAH5", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusif
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Println("Declare is error : ", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, //routing key
		false,  //mandatory
		false,  // immediati
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte("sapitenglah"),
		})
	if err != nil {
		log.Println("Publish is error : ", err)
	}
}
