package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var manager = NewClientManager()

func HandleWS(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	if clientID == "" {
		http.Error(w, "client_id is required", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	manager.Add(clientID, conn)
	client, _ := manager.Get(clientID)

	log.Printf("Client connected: %s\n", clientID)

	defer func() {
		manager.Remove(clientID)
		conn.Close()
		log.Printf("Client disconnected: %s\n", clientID)
	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var cmd Command
		if err := json.Unmarshal(data, &cmd); err != nil {
			continue
		}

		// AUTH
		if !authRequired(client, cmd.Type) {
			conn.WriteJSON(map[string]any{
				"type": "error",
				"data": "unauthorized",
			})
			continue
		}

		switch cmd.Type {

		case "ping":
			handlePing(conn)

		case "auth":
			handleAuth(client, cmd.Data)

		case "message":
			handleMessage(client, cmd.Data)
		}
	}
}
