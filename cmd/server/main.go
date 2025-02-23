package main

import (
	"log"
	"simple_websocket/internal/app"
	"simple_websocket/internal/config"
	"simple_websocket/internal/controller"
	"simple_websocket/pkg/logger"

	"github.com/gin-gonic/gin"
)

// main is the entry point of the application.
func main() {
	// Load application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("⚠️ Failed to load configuration: %s", err)
	}

	// Initialize logger with loaded config
	logger.InitLogger(cfg.Logger)

	// Log successful initialization
	logger.WriteLog.Info("✅ Logger initialized successfully!")

	hub := app.NewHub()
	go hub.Run()

	r := gin.Default()
	r.GET("/ws", controller.WebSocketHandler(hub))

	logger.WriteLog.Info("✅ Server started on port 3000")
	r.Run(":3000")
}
