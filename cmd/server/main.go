package main

import (
	"fmt"
	"log"
	"simple_websocket/internal/config"
	"simple_websocket/pkg/logger"
)

// main intializes the application (starting point of the application)
func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("failed to load configuration from the file. ERROR: %s", err)
		return
	}

	fmt.Println(cfg)

	logger.InitLogger(cfg.Logger)

	logger.WriteLog.Error("Hello")
	logger.WriteLog.Debug("Hello")
	logger.WriteLog.Info("Hello")
	logger.WriteLog.Warn("Hello")
}
