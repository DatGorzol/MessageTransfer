package websocket

func authRequired(client *Client, cmdType string) bool {
	// auth и ping разрешены всем
	if cmdType == "auth" || cmdType == "ping" {
		return true
	}

	return client.Auth
}
