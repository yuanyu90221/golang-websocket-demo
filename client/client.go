package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello from client!\n"))
	if err != nil {
		log.Fatal(err)
		return
	}
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Received: %s\n", string(message))
}
