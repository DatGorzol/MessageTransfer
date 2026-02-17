package websocket

import "sync"

type OfflineStore struct {
	messages map[string][]any // userID -> messages
	mu       sync.Mutex
}

func NewOfflineStore() *OfflineStore {
	return &OfflineStore{
		messages: make(map[string][]any),
	}
}

func (o *OfflineStore) Add(userID string, msg any) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.messages[userID] = append(o.messages[userID], msg)
}

func (o *OfflineStore) PopAll(userID string) []any {
	o.mu.Lock()
	defer o.mu.Unlock()

	msgs := o.messages[userID]
	delete(o.messages, userID)

	return msgs
}

var offlineStore = NewOfflineStore()
