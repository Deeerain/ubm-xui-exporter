package main

import (
	"log/slog"
	"os"
)

func InitLogger(level string) {
	var loggerLevel slog.Level
	switch level {
	case "debug":
		loggerLevel = slog.LevelDebug
	case "warn":
		loggerLevel = slog.LevelWarn
	case "error":
		loggerLevel = slog.LevelError
	default:
		loggerLevel = slog.LevelInfo
	}

	loggerOpts := &slog.HandlerOptions{
		Level: loggerLevel,
	}
	logger := slog.NewTextHandler(os.Stdout, loggerOpts)
	slog.SetDefault(slog.New(logger))
}
