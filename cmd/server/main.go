package main

import (
	"log"
	"simple_websocket/internal/config"
	"simple_websocket/pkg/logger"
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
}
