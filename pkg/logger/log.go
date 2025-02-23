// Package logger provides a structured logging utility using Zap.
// It supports console and file logging with optional log rotation.
package logger

import (
	"os"
	"simple_websocket/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// WriteLog is the exported logger instance for application-wide logging.
var WriteLog *zap.Logger

// LogLevel represents different logging levels.
type LogLevel int

const (
	// Define log levels explicitly
	LogDisabled LogLevel = iota
	LogFatal
	LogPanic
	LogDPanic
	LogError
	LogWarn
	LogInfo
	LogDebug
)

// logLevelMap maps LogLevel constants to their corresponding Zap log levels.
var logLevelMap = map[LogLevel]zapcore.Level{
	LogFatal:  zapcore.FatalLevel,
	LogPanic:  zapcore.PanicLevel,
	LogDPanic: zapcore.DPanicLevel,
	LogError:  zapcore.ErrorLevel,
	LogWarn:   zapcore.WarnLevel,
	LogInfo:   zapcore.InfoLevel,
	LogDebug:  zapcore.DebugLevel,
}

// parseLogLevel converts an integer to a valid Zap log level.
func parseLogLevel(level int) zapcore.Level {
	if logLevel, exists := logLevelMap[LogLevel(level)]; exists {
		return logLevel
	}
	return zapcore.ErrorLevel // Default log level
}

// InitLogger initializes the logging system.
// It supports console logging, file logging, and log rotation.
func InitLogger(logConfig config.LoggerConfig) *zap.Logger {

	logLevel := parseLogLevel(logConfig.Level)

	// Configure encoder settings
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.StacktraceKey = ""
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	var cores []zapcore.Core

	// File logging configuration
	if logConfig.EnableFile {
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logConfig.Filename,
			MaxSize:    200, // MB
			MaxBackups: 5,
			MaxAge:     7, // Days
			Compress:   true,
		})

		cores = append(cores, zapcore.NewCore(fileEncoder, writer, logLevel))
	}

	// Console logging configuration
	if logConfig.EnableConsole {
		cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel))
	}

	// If no logging is enabled, return a no-op logger
	if len(cores) == 0 {
		WriteLog = zap.NewNop() // No-op logger (disables logging)
		return WriteLog
	}

	// Combine all logging cores
	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller())

	WriteLog = logger // Assign global logger
	return logger
}
