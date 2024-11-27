package ws

import (
	"log/slog"
	"sync"
)

type GameManager struct {
	DuelQueue   []*Client
	SoloPlayers []*Client

	mu sync.Mutex
}

func NewGameManager() *GameManager {
	return &GameManager{
		DuelQueue:   make([]*Client, 0),
		SoloPlayers: make([]*Client, 0),
	}
}

// Duel with friend and duel with random user
// Duel with friend will be provided a GameID(Link)
func (g *GameManager) HandleDuelGame(c *Client) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if len(g.DuelQueue) > 0 {
		// Check for potential bug
		for _, queuedClient := range g.DuelQueue {
			if queuedClient == c {
				slog.Error("Client already in queue", "Username", c.UserState.Username)
				return
			}
		}
		opponent := g.DuelQueue[0]    // first user go to Queue
		g.DuelQueue = g.DuelQueue[1:] // remove opponent
		duel := NewDuel(opponent, c)
		go duel.Start()
		slog.Info("GAME HAS STARTED")
	} else {
		g.DuelQueue = append(g.DuelQueue, c)
		msg := &ResponseMessage{
			Message: "Queueing",
			State:   QUEUEING,
			Score:   c.UserState.Score,
		}

		c.ReceiveMessage(msg)
	}
}

func (g *GameManager) HandleSoloGame(c *Client) {

}
