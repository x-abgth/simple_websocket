// Package app handles WebSocket connection logic.
package app

import (
	"time"

	"simple_websocket/internal/model"
	"simple_websocket/pkg/logger"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// HandleWebSocket manages a new WebSocket connection.
func HandleWebSocket(hub *Hub, conn *websocket.Conn) {
	client := &Client{Conn: conn, Send: make(chan model.Message)}
	hub.Register <- client

	go client.ReadMessages(hub)
	go client.WriteMessages()
}

// ReadMessages handles incoming messages from a client.
func (c *Client) ReadMessages(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var msg model.Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			logger.WriteLog.Error("Error reading message:", zap.Error(err))
			break
		}
		msg.Timestamp = time.Now()
		hub.Broadcast <- msg

		// Log message received
		logger.WriteLog.Info("📩 New message received",
			zap.String("sender", msg.Sender),
			zap.String("content", msg.Content),
			zap.Time("timestamp", msg.Timestamp),
		)
	}
}

// WriteMessages sends messages from the server to the client.
func (c *Client) WriteMessages() {
	defer c.Conn.Close()

	for msg := range c.Send {
		err := c.Conn.WriteJSON(msg)
		if err != nil {
			logger.WriteLog.Error("Error sending message:", zap.Error(err))
			break
		}
	}
}
