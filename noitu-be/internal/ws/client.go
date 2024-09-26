package ws

import (
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/slog"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Client struct {
	Hub *Hub

	Conn *websocket.Conn

	Resp      chan *ResponseMessage
	Req       chan *ClientMessage
	UserState User
}

func (c *Client) ReadMessage() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		var message ClientMessage
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error("error: ", slog.Any("ReadJson", err))
			}
			slog.Error("Err", slog.Any("err:", err))
			break
		}
		slog.Info("message", slog.Any("message", message))
		c.Req <- &message
	}
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Hub.Unregister <- c
	}()
	for {
		message, ok := <-c.Resp
		if !ok {
			return
		}
		c.Conn.WriteJSON(message)
	}
}
