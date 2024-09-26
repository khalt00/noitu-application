package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/khalt00/noitu/internal/ws"
)

type GameController interface {
	Register(*gin.Context)
}

type gameController struct {
	hub *ws.Hub
	gm  *ws.GameManager
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewGameController(hub *ws.Hub, gm *ws.GameManager) GameController {
	return &gameController{
		hub: hub,
		gm:  gm,
	}
}

// When a user connect to our website
// will have a box to input username, we will find base on the username only

type RegisterStruct struct {
	Username string      `form:"username"`
	Category ws.GameType `form:"category"`
}

func (c *gameController) Register(ctx *gin.Context) {

	var req RegisterStruct
	if err := ctx.ShouldBind(&req); err != nil {
		slog.Error("invalid req", "ERR", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		slog.Error("invalid upgrade websocket", "err", err)
		return
	}

	client := &ws.Client{
		Hub:  c.hub,
		Conn: conn,
		Resp: make(chan *ws.ResponseMessage, 1024),
		Req:  make(chan *ws.ClientMessage, 1024),
		UserState: ws.User{
			Username:  req.Username,
			Category:  req.Category,
			IsPlaying: true,
			Score:     500,
		},
	}

	client.Hub.Register <- client

	go client.ReadMessage()
	go client.WriteMessage()
	switch req.Category {
	case ws.DuelGame:
		c.gm.HandleDuelGame(client)
	case ws.Alone:
		c.gm.HandleSoloGame(client)
	}
}
