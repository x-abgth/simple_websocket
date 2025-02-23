// Package config handles application configuration loading from .env using Viper.
// It provides structured configuration management for the application, ensuring
// all settings are properly loaded and accessible.
package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all application configuration values.
// This struct is the root-level configuration used across the application.
type Config struct {
	Logger LoggerConfig // Logger configuration
}

// LoggerConfig holds the logging configuration settings.
type LoggerConfig struct {
	Filename      string // File path for storing logs
	EnableFile    bool   // Enable logging to a file
	EnableConsole bool   // Enable logging to console (stdout)
	Level         int    // Log level (0=FATAL, 1=PANIC, 2=ERROR, etc.)
}

// LoadConfig initializes Viper and loads the application configuration from a .env file.
// It automatically reads environment variables if the .env file is not found.
func LoadConfig() (*Config, error) {

	viper.SetConfigFile(".env")                            // Define the config file
	viper.SetConfigType("env")                             // Specify file type as .env
	viper.AutomaticEnv()                                   // Automatically read system environment variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Allow nested env keys like LOG.LEVEL

	// Attempt to read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("⚠️ Warning: No .env file found. Using system environment variables. Error: %s", err.Error())
	}

	// Return the fully constructed Config object
	return &Config{Logger: readLoggerConfig()}, nil
}

// readLoggerConfig loads and returns logging configuration from environment variables.
// It ensures that all necessary log settings are properly fetched and structured.
func readLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Filename:      viper.GetString("LOG_FILE"),         // Log file path
		EnableFile:    viper.GetBool("ENABLE_FILE_LOG"),    // Enable file logging
		EnableConsole: viper.GetBool("ENABLE_CONSOLE_LOG"), // Enable console logging
		Level:         viper.GetInt("LOG_LEVEL"),           // Logging level
	}
}
