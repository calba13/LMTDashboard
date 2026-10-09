package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"os"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", handleWebSocketConnection)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to websocket:", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received: %s", message)
		// Handle incoming messages here
	}
}
