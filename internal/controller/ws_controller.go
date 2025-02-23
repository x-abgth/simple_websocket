// Package controller contains HTTP and WebSocket route handlers.
package controller

import (
	"net/http"

	"simple_websocket/internal/app"
	"simple_websocket/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all origins
}

// WebSocketHandler handles WebSocket connection requests.
func WebSocketHandler(hub *app.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			logger.WriteLog.Error("failed to upgrade HTTP to WS", zap.Error(err))
			return
		}

		app.HandleWebSocket(hub, conn)
	}
}
