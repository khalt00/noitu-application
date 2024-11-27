package ws

import (
	"sync"
)

type Hub struct {
	mu      sync.Mutex
	Clients map[*Client]interface{}

	Register chan *Client

	Unregister chan *Client

	Message     chan *ClientMessage
	GameManager *GameManager
}

func NewHub(gm *GameManager) *Hub {
	return &Hub{
		Clients:     make(map[*Client]interface{}),
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		GameManager: gm,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client] = true
			h.mu.Unlock()
		case client := <-h.Unregister:
			h.mu.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
			}
			h.mu.Unlock()
		}
	}
}
