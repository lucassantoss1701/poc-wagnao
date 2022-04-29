package main

import (
	"fmt"
	"log"
	"net/http"
	"subscriber/framework/queue"

	"github.com/streadway/amqp"
)

func main() {
	messageChannel := make(chan amqp.Delivery)

	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()
	defer ch.Close()

	rabbitMQ.Consume(messageChannel)

	go func() {
		for {
			select {
			case msg := <-messageChannel:
				fmt.Println("MENSANGEM DA FILA: ", string(msg.Body))
			}
		}
	}()

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
