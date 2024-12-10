package logger

import (
	"log/slog"
	"os"
)

const logFile = "logs/client.log"

var Logger = Make()

func Make() *slog.Logger {
	file, _ := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND, 0644)
	logger := slog.New(slog.NewTextHandler(file, nil))
	return logger
}

// TODO добавить finish
