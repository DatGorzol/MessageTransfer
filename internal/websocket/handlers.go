package websocket

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

func handlePing(conn *websocket.Conn) {
	resp := map[string]any{
		"type": "pong",
		"data": map[string]any{
			"timestamp": time.Now().Unix(),
		},
	}

	conn.WriteJSON(resp)
}

func handleMessage(fromClient *Client, data []byte) {
	var payload MessagePayload
	if err := json.Unmarshal(data, &payload); err != nil {
		fromClient.Send(map[string]any{"type": "error", "data": "bad payload"})
		return
	}

	resp := map[string]any{
		"type": "message",
		"data": map[string]string{
			"from": fromClient.UserID,
			"text": payload.Text,
		},
	}

	dispatcher.Dispatch(Envelope{
		From: fromClient,
		To:   payload.To,
		Data: resp,
	})
}

func handleAuth(client *Client, data []byte) {
	var payload AuthPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		client.Send(map[string]any{"type": "auth", "data": "invalid"})
		return
	}

	if payload.Token != "valid-token" {
		client.Send(map[string]any{"type": "auth", "data": "failed"})
		return
	}

	client.Auth = true
	client.UserID = client.ID

	client.Send(map[string]any{
		"type": "auth",
		"data": "success",
	})

	msgs := offlineStore.PopAll(client.UserID) //доставка офланй
	for _, msg := range msgs {
		client.Send(msg)
	}
}
