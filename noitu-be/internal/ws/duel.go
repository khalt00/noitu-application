package ws

import (
	"log/slog"
	"time"

	"github.com/khalt00/noitu/internal/dict"
	"github.com/khalt00/noitu/pkg/utils"
)

type Duel struct {
	GameID string
	User1  *Client
	User2  *Client

	CurrentWord string
	GameTime    uint
	CurrentTurn *Client

	quit chan struct{}
}

func NewDuel(user1, user2 *Client) *Duel {
	return &Duel{
		User1: user1,
		User2: user2,
		quit:  make(chan struct{}),
	}
}

func (d *Duel) Start() {
	d.User1.Resp <- &ResponseMessage{
		Message:   "Hi from game",
		IsPlaying: true,
	}

	d.User2.Resp <- &ResponseMessage{
		Message:   "Hi from game",
		IsPlaying: true,
	}
	d.CurrentWord = dict.GetRandomWord()
	d.CurrentTurn = d.User1
	d.sendWord()
	go d.gameLoop()
}

func (d *Duel) gameLoop() {
	d.CurrentTurn = d.User1
	gameOver := false

gameloop:
	for !gameOver {
		select {
		case message := <-d.CurrentTurn.Req:
			if message == nil {
				return
			}
			d.HandleCheckCorrectWord(message.Word)
		case <-time.After(5 * time.Second):
		case <-d.quit:
			gameOver = true
			break gameloop
		}
	}
}

func (d *Duel) sendWord() {
	// send message to current user have to answer
	d.CurrentTurn.Resp <- &ResponseMessage{
		Message:   utils.CombineString(d.CurrentWord),
		IsPlaying: true,
	}

	// send message to another user wait for the answer
	another := d.getAnotherUser()
	msg := "Wait for another user"
	another.Resp <- &ResponseMessage{
		Message:   utils.CombineString(msg),
		IsPlaying: true,
	}
}

func (d *Duel) getAnotherUser() (anotherUser *Client) {
	if d.CurrentTurn == d.User1 {
		return d.User2
	} else {
		return d.User1
	}
}

// Correct : + 5 + streak
// Incorrect : - 5, streak to 0
// Streak ++ after win each game
// not correct answer
func (d *Duel) Score() {}

func (d *Duel) EndGame() {
	d.User1.Resp <- &ResponseMessage{
		Message:   "Game over. Choose an option: 1) Play Again 2) Quit",
		GameOver:  true,
		IsPlaying: false,
	}
	d.User2.Resp <- &ResponseMessage{
		Message:   "Game over. Choose an option: 1) Play Again 2) Quit",
		IsPlaying: false,
		GameOver:  true,
	}

	close(d.quit)
	go d.handleEndGameOptions(d.User1)
	go d.handleEndGameOptions(d.User2)

}

func (d *Duel) handleEndGameOptions(user *Client) {
	message := <-user.Req
	if message == nil {
		return
	}
	switch message.Word {
	case string(PlayAgain):
		user.Hub.GameManager.HandleDuelGame(user)
	case string(Quit):
		user.Resp <- &ResponseMessage{
			Message:   "Thanks for playing!",
			IsPlaying: false,
		}
		return
	}
}

// TODO:
// - When out of time, => - score of currentTurn, + score for other
// - both will receive message keep playing or quit
func (d *Duel) HandleCheckCorrectWord(word string) {
	correct := utils.CompareCorrectConnectWord(d.CurrentWord, word)
	slog.Info("correct?", "isCorrect", correct)
	if !correct {
		d.EndGame()
	}

	if d.CurrentTurn == d.User1 {
		d.CurrentTurn = d.User2
	} else {
		d.CurrentTurn = d.User1
	}

	d.CurrentWord = word
	d.sendWord()
}
