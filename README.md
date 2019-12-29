# Golang with RabbitMQ 

This repo for example how to implement Golang messaging with RabbitMQ

##  How to run
- you can install rabbitMQ in your computer tutorial:
  - if you mac user  [link](https://www.rabbitmq.com/install-homebrew.html)
  - if you linux user [link](https://www.rabbitmq.com/install-debian.html)
  - if you windows user [link](https://www.rabbitmq.com/install-windows.html) 
- install dependencies go get
    - `go get github.com/streadway/amqp`

## Fitur
- Receiver
  - this fitur for receive message from **sender**
- Sender
  - this fitur for sending message to **receiver**

## FYI 
- Full tutorial in this repo you can read [this](https://medium.com/@cakazies/messaging-golang-with-rabbitmq-2ed1ccf8314)

## Run Sender
`go run send/main.go`

## Run Receiver
`go run receive/main.go`
