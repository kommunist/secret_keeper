package logger

import (
	"log/slog"
	"os"
)

const logFile = "logs/client.log"

var Logger = Make()

// Данный логгер переопределяет основной вывод в файл, так как терминал занят отображением
func Make() *slog.Logger {
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err) // Кажется, паниковать в данном случае допустимо. Тем более логгер глобальный
	}
	logger := slog.New(slog.NewTextHandler(file, nil))
	return logger
}
