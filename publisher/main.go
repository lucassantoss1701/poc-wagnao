package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"publisher/queue"
	"time"
)

type Message struct {
	ID      int
	Message string
}

func main() {
	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()
	defer ch.Close()

	time.Sleep(10 * time.Second)
	for i := 0; i < 100; i++ {

		message := &Message{
			ID:      i,
			Message: fmt.Sprintf("Mensagem %d", i),
		}
		msg, _ := json.Marshal(message)
		rabbitMQ.Notify(string(msg), "application/json", "amq.direct", "new-message")
	}

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
