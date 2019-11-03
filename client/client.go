package client

import (
	"fmt"

	"github.com/streadway/amqp"
)

type client struct {
	connectionURL string
	connection    *amqp.Connection
	Channel       *amqp.Channel
}

func NewClient(rmqURL string, user string, pass string) *client {

	connectionURL := fmt.Sprintf("amqp://%s:%s@%s", user, pass, rmqURL)

	fmt.Println("Connecting to: " + connectionURL)

	connection, err := amqp.Dial(connectionURL)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	c := client{
		connectionURL: connectionURL,
		connection:    connection,
		Channel:       channel,
	}

	return &c
}
