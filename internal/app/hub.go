// Package app contains core application logic, including WebSocket handling.
package app

import (
	"sync"

	"simple_websocket/internal/model"
	"simple_websocket/pkg/logger"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Client represents a connected WebSocket user.
type Client struct {
	Conn *websocket.Conn    // WebSocket connection
	Send chan model.Message // Channel for sending messages
}

// Hub manages all connected clients and message broadcasting.
type Hub struct {
	Clients    map[*Client]bool   // Set of connected clients
	Broadcast  chan model.Message // Channel for broadcasting messages
	Register   chan *Client       // Clients joining
	Unregister chan *Client       // Clients leaving
	mu         sync.Mutex         // Ensures thread safety
}

// NewHub creates a new instance of the Hub.
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan model.Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run starts the Hub event loop to manage client connections.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client] = true
			h.mu.Unlock()
			logger.WriteLog.Info("👨🏻✅ New Client connection established!")

		case client := <-h.Unregister:
			h.mu.Lock()
			if _, exists := h.Clients[client]; exists {
				delete(h.Clients, client)
				close(client.Send)
				client.Conn.Close()
				logger.WriteLog.Info("👨🏻❌ Client disconnected!")
			}
			h.mu.Unlock()

		case message := <-h.Broadcast:
			h.mu.Lock()

			logger.WriteLog.Info("📢 Broadcasting message",
				zap.String("sender", message.Sender),
				zap.String("content", message.Content),
				zap.Time("timestamp", message.Timestamp),
			)

			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					// removes unresponsive clients / disconnected clients
					delete(h.Clients, client)
					close(client.Send)
					client.Conn.Close()
					logger.WriteLog.Info("Removed unresponsive client!")
				}
			}
			h.mu.Unlock()
		}
	}
}
