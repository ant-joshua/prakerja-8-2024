package main

import "time"

type ChatMessage struct {
	Username string
	Message  string
}

var chat = make(chan ChatMessage)

func main() {

	go func() {
		chat <- ChatMessage{Username: "doe", Message: "Hello, Go Routine!"}
	}()

	for {
		select {
		case msg := <-chat:
			println(msg.Username, ":", msg.Message)
		}

		go func() {
			time.Sleep(1 * time.Second)
			chat <- ChatMessage{Username: "john", Message: "Hello, World!"}

		}()

		go func() {
			time.Sleep(1 * time.Second)
			chat <- ChatMessage{Username: "doe", Message: "Hello, Go Routine!"}

		}()
	}

}
