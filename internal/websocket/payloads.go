package websocket

type MessagePayload struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

type PingPayload struct {
	Timestamp int64 `json:"timestamp"`
}

type AuthPayload struct {
	Token string `json:"token"`
}
