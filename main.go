package main

import (
	"allopopot-email-service/config"
	"allopopot-email-service/queues"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitExchange() *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", config.AMQP_USERNAME, config.AMQP_PASSWORD, config.AMQP_HOST, config.AMQP_PORT))
	if err != nil {
		log.Panicln("Failed to connect to AMQP server.")
	}
	log.Println("Connected to AMQP server.")
	// defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panicln("Failed to init channel.")
	}
	log.Println("Channel initialized.")
	// defer ch.Close()

	err = ch.ExchangeDeclare(config.AMQP_EXCHANGE_NAME, "fanout", true, false, false, false, nil)
	if err != nil {
		log.Panicln("Failed to declare exchange.")
	}
	log.Println("Exchange declared.")

	return ch
}

func main() {
	channel := InitExchange()
	queues.InitEmailDispatcherQueue(channel)
	defer channel.Close()
}
