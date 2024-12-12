package main

import (
	"secret_keeper/internal/client/app"
	"secret_keeper/internal/client/logger"
)

func main() {
	a := app.Make()

	a.Start()
	logger.Logger.Info("ququququ")
}
