package main

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to connect to rabbitmq
	rabbitConn, err := connectRabbitMQ()
	if err != nil {
		log.Println("Failed to connect to RabbitMQ: ", err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	log.Println("Connected to RabbitMQ!")

	// start listening for messages

	// create consumer

	// watch the queue and consume events
}

// connectRabbitMQ
func connectRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
