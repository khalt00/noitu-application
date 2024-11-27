package ws

type GameType string

const (
	DuelGame   GameType = "duel"
	Alone      GameType = "alone"
	WithFriend GameType = "w_friend"
)

type GameState string

const (
	NONE                 GameState = "NONE"
	QUEUEING             GameState = "QUEUEING"
	PLAYING              GameState = "PLAYING"
	WAIT_FOR_ANOTHERUSER GameState = "WAITING"
	ENDING               GameState = "ENDING"
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
	Message     string    `json:"msg"`
	CurrentWord string    `json:"currentWord"`
	State       GameState `json:"state"`
	Score       uint64    `json:"score"`
}

type EndGameOptions string

const (
	PlayAgain EndGameOptions = "play_again"
	Quit      EndGameOptions = "quit"
)
