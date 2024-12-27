package logger

import (
	"log/slog"
	"os"
)

const logFile = "logs/client.log"

var Logger = Make()

// Данный логгер переопределяет основной вывод в файл, так как терминал занят отображением
func Make() *slog.Logger {
	file, _ := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND, 0644)
	logger := slog.New(slog.NewTextHandler(file, nil))
	return logger
}
