package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/a.alepis/rabbit_logger/client"
)

func main() {
	fmt.Println("Starting rabbit logger...")

	queuePtr := flag.String("queue", "noqueue", "The name of the queue to log")
	urlPtr := flag.String("url", "0.0.0.0:5672", "RabbitMq URL:PORT")
	usernamePtr := flag.String("username", "guest", "RabbitMq username")
	passwordPtr := flag.String("password", "guest", "RabbitMq password")

	flag.Parse()

	client := client.NewClient(*urlPtr, *usernamePtr, *passwordPtr)

	msgs, err := client.Channel.Consume(*queuePtr, "", false, false, false, false, nil)

	if err == nil {
		fmt.Println("Consuming from queue: " + *queuePtr)
	} else {
		panic("Could not consume from queue: " + *queuePtr)
	}

	for msg := range msgs {
		log := fmt.Sprintf("[%s] %s", time.Now().Format(time.UnixDate), string(msg.Body))

		fmt.Println(log)

		msg.Ack(false)
	}

}
