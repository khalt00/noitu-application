package ws

import "golang.org/x/exp/slog"

type GameManager struct {
	DuelQueue []*Client

	SoloPlayers []*Client
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
	slog.Info("len duels", "len", len(g.DuelQueue))
	if len(g.DuelQueue) > 0 {
		opponent := g.DuelQueue[0]    // first user go to Queue
		g.DuelQueue = g.DuelQueue[1:] // remove opponent
		duel := NewDuel(opponent, c)
		go duel.Start()
	} else {
		g.DuelQueue = append(g.DuelQueue, c)
	}
}

func (g *GameManager) HandleSoloGame(c *Client) {

}
