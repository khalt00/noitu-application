package ws

import (
	// "fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/khalt00/noitu/internal/dict"
	"github.com/khalt00/noitu/pkg/utils"
)

type Duel struct {
	GameID string
	User1  *Client
	User2  *Client

	mu          sync.Mutex
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
	msg := &ResponseMessage{
		Message: "Hi from game",
	}
	msg2 := &ResponseMessage{
		Message: "Hi from game",
		State:   WAIT_FOR_ANOTHERUSER,
	}

	d.User1.ReceiveMessage(msg)
	d.User2.ReceiveMessage(msg2)

	d.mu.Lock()
	d.CurrentWord = dict.GetRandomWord()
	d.CurrentTurn = d.User1
	d.mu.Unlock()
	d.sendWord()
	go d.gameLoop()
}

func (d *Duel) gameLoop() {
	gameOver := false

gameloop:
	for !gameOver {
		slog.Info("current user", "username", d.CurrentTurn.UserState.Username, "word", d.CurrentWord)
		select {
		case message := <-d.CurrentTurn.Req:
			slog.Info("wtf is going on?", "req", message)
			if message == nil {
				return
			}
			d.HandleCheckCorrectWord(message.Word)
		// case _, ok := <-d.User1.Req:
		// 	slog.Info("go here??? wtf", "name", d.CurrentTurn.UserState.Username)
		// 	if !ok {
		// 		fmt.Println("user1 disconnected")
		// 		gameOver = true
		// 		break gameloop
		// 	}
		// case _, ok2 := <-d.User2.Req:
		// 	slog.Info("go here???2 wtf", "name", d.CurrentTurn.UserState.Username)
		// 	if !ok2 {
		// 		fmt.Println("user2 disconnected")
		// 		gameOver = true
		// 		break gameloop
		// 	}
		case <-time.After(2 * time.Second):
			slog.Warn("No message received in time for current turn", "username", d.CurrentTurn.UserState.Username)
		case <-d.quit:
			gameOver = true
			break gameloop
		}
	}
}

func (d *Duel) sendWord() {
	// send message to current user have to answer
	msg1 := &ResponseMessage{
		Message:     utils.CombineString("Your turn"),
		CurrentWord: d.CurrentWord,
		State:       PLAYING,
	}

	another := d.getAnotherUser()
	msg2 := &ResponseMessage{
		Message:     utils.CombineString("Wait for another user"),
		CurrentWord: d.CurrentWord,
		State:       WAIT_FOR_ANOTHERUSER,
	}
	d.CurrentTurn.ReceiveMessage(msg1)
	another.ReceiveMessage(msg2)

}

// hehe
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
// This function only run when user answer incorrect
// So the current User will be -, another will be +
func (d *Duel) HandleScore() {
	// What if a user < 1 (negative number, this one suck)
	if d.CurrentTurn.UserState.Score > 1 {
		d.CurrentTurn.UserState.Score = d.CurrentTurn.UserState.Score - 5
	}
	another := d.getAnotherUser()
	another.UserState.Score = another.UserState.Score + 5

}

func (d *Duel) EndGame() {
	msg1 := &ResponseMessage{
		Message: "Game over. Choose an option: 1) Play Again 2) Quit",
		State:   ENDING,
	}
	msg2 := &ResponseMessage{
		Message: "Game over. Choose an option: 1) Play Again 2) Quit",
		State:   ENDING,
	}
	d.CurrentTurn.ReceiveMessage(msg1)
	another := d.getAnotherUser()
	another.ReceiveMessage(msg2)

	close(d.quit)
	go d.handleEndGameOptions(d.CurrentTurn)
	go d.handleEndGameOptions(another)
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
			Message: "Thanks for playing!",
			State:   NONE,
		}
		user.Quit()
		return
	default:
		user.Quit()
		return
	}
}

// TODO:
// - When out of time, => - score of currentTurn, + score for other
// - both will receive message keep playing or quit
func (d *Duel) HandleCheckCorrectWord(word string) {
	correct := utils.CompareCorrectConnectWord(d.CurrentWord, word)
	slog.Info("correct", "info", correct)
	if !correct {
		d.HandleScore()
		d.EndGame()
		return
	}

	if d.CurrentTurn == d.User1 {
		d.CurrentTurn = d.User2
	} else {
		d.CurrentTurn = d.User1
	}

	d.CurrentWord = word
	d.sendWord()
}
