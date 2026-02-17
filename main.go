package main

import (
	"log"
	"net/http"

	ws "MessageTransfer/internal/websocket"
)

func main() {
	ws.StartDispatcher()

	http.HandleFunc("/ws", ws.HandleWS)

	log.Println("WebSocket server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
