package ws

type GameType string

const (
	DuelGame   GameType = "duel"
	Alone      GameType = "alone"
	WithFriend GameType = "w_friend"
)

type GameState string

const (
	Playing  GameState = "Playing"
	Queueing GameState = "Queueing"
)

type ClientMessage struct {
	Word string `json:"word"`
}

type User struct {
	Username  string
	Score     uint64
	Streak    uint8
	IsPlaying bool
	Category  GameType
}

type ResponseMessage struct {
	Message   string      `json:"msg"`
	IsPlaying bool        `json:"isPlaying"`
	GameOver  bool        `json:"gameOver"`
	Data      interface{} `json:"data"`
}

type EndGameOptions string

const (
	PlayAgain EndGameOptions = "play_again"
	Quit      EndGameOptions = "quit"
)
