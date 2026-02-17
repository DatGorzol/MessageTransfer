package websocket

type Envelope struct {
	From *Client
	To   string
	Data any
}
