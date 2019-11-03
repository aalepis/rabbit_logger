package main

import (
	"fmt"
	"os"
	"time"

	"github.com/a.alepis/rabbit_logger/client"
)

func main() {
	fmt.Println("Starting rabbit logger...")

	client := client.NewClient(os.Args[2], os.Args[3], os.Args[4])

	msgs, _ := client.Channel.Consume(os.Args[1], "", false, false, false, false, nil)

	for msg := range msgs {
		log := fmt.Sprintf("[%s] %s", time.Now().Format(time.UnixDate), string(msg.Body))

		fmt.Println(log)

		msg.Ack(false)
	}

}
