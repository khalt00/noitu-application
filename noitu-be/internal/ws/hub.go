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
		// Message:    make(chan *ClientMessage),
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
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Req)
				close(client.Resp)
			}
			// case message := <-h.Message:
			// switch message.Category {
			// case DuelGame:
			// 	slog.Info("user one v one")
			// 	if len(h.Queueing) > 0 {
			// 		opponent := h.Queueing[0]   // first user go to Queue
			// 		h.Queueing = h.Queueing[1:] // remove opponent
			// 		// TODO: GameStarted
			// 		duel := NewDuel(message.Client, opponent)
			// 		go duel.Start()
			// 	} else {
			// 		h.Queueing = append(h.Queueing, message.Client)
			// 	}
			// }
		}
	}
}
