// Package model defines the data structures used in the application.
package model

import "time"

// Message represents a chat message sent over WebSockets.
type Message struct {
	Sender    string    `json:"sender"`    // Name of the sender
	Content   string    `json:"content"`   // Message text
	Timestamp time.Time `json:"timestamp"` // Time when message was sent
}
