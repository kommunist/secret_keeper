package main

import (
	"os"
	"secret_keeper/internal/client/app"
	"secret_keeper/internal/client/logger"
)

func main() {
	a, err := app.Make()
	if err != nil {
		logger.Logger.Error("Error when make app", "err", err)
		os.Exit(1)
	}

	a.Start()
}
