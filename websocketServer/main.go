package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var c chan string

func handleWebSocket(conn *websocket.Conn) {
	defer conn.Close()

	for {
		// Read message from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		// Chuck the message into a channel
		c <- string(msg)
		// Print received message
		returnMessage := fmt.Sprintf("Received message: %s", msg)

		// Echo the message back to the client
		err = conn.WriteMessage(websocket.TextMessage, []byte(returnMessage))
		if err != nil {
			log.Println("Error writing message:", err)
			return
		}
	}
}

// Handle the message and don't worry too much about time since its concurrent
func handleMessage() {
	for {
		fmt.Printf("Message: %s\n", <-c)
	}
}

func main() {
	// WebSocket endpoint
	endpoint := "/ws"

	c = make(chan string, 100)
	go handleMessage()
	// Configure WebSocket route
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error upgrading to WebSocket:", err)
			return
		}
		handleWebSocket(conn)
	})

	// Start WebSocket server
	serverAddr := "127.0.0.1:8080"
	log.Printf("WebSocket server listening on ws://%s%s\n", serverAddr, endpoint)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatal("Error starting WebSocket server:", err)
	}
}
