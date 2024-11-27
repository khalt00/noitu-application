package ws

import (
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"log/slog"
)

type Client struct {
	Hub       *Hub
	Conn      *websocket.Conn
	Resp      chan *ResponseMessage
	Req       chan *ClientMessage
	UserState User

	quitOnce sync.Once
}

func (c *Client) ReadMessage() {
	defer c.Quit()
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
		select {
		case c.Req <- &message:
			slog.Info("Message received and sent to Req", c.UserState.Username, slog.Any("message", message))
		case <-time.After(2 * time.Second): // Timeout if Req isn't consumed
			slog.Warn("Dropping message due to timeout", slog.Any("message", message))
		}
	}
}

func (c *Client) WriteMessage() {
	defer c.Quit()
	for {
		message, ok := <-c.Resp
		if !ok {
			return
		}
		c.Conn.WriteJSON(message)
	}
}

func (c *Client) Quit() {
	c.quitOnce.Do(func() {
		// Unregister client from the hub
		c.Hub.Unregister <- c
		// Close the WebSocket connection
		c.Conn.Close()
		// Close request and response channels
		close(c.Req)
		close(c.Resp)
	})
}

func (c *Client) ReceiveMessage(msg *ResponseMessage) error {
	if c == nil {
		slog.Error("Client does not exists", "Err", nil)
		return errors.New("client does not exists")
	}
	if msg == nil {
		slog.Error("Received nil message", "Username", c.UserState.Username)
		return errors.New("invalid nil message")
	}
	if c.Resp == nil {
		slog.Error("Response channel is nil", "Username", c.UserState.Username)
		return errors.New("nil response")
	}
	msg.Score = c.UserState.Score
	select {
	case c.Resp <- msg:
		// Message sent successfully
	default:
		slog.Warn("Response channel is full, message dropped", "Username", c.UserState.Username)
	}

	return nil
}
