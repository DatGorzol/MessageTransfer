package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	Conn   *websocket.Conn
	Auth   bool
	UserID string
	Role   string
}

func (c *Client) Send(v any) error {
	return c.Conn.WriteJSON(v)
}

type ClientManager struct {
	clients map[string]*Client
	mu      sync.RWMutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[string]*Client),
	}
}

func (cm *ClientManager) Add(id string, conn *websocket.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.clients[id] = &Client{
		ID:   id,
		Conn: conn,
		Auth: false,
	}
}

func (cm *ClientManager) Remove(id string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.clients, id)
}

func (cm *ClientManager) Get(id string) (*Client, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	client, ok := cm.clients[id]
	return client, ok
}
