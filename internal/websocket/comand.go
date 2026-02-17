package websocket

import "encoding/json"

type Command struct {
	Type string          `json:"type"` // message | ping | auth
	Data json.RawMessage `json:"data"`
}
