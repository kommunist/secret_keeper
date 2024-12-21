package main

import (
	"os"
	"secret_keeper/internal/client/app"
	"secret_keeper/internal/client/logger"
)

var buildVersion string = "N/A"
var buildDate string = "N/A"

func main() {
	logger.Logger.Info("======Started======")
	logger.Logger.Info("Build info", "version", buildVersion)
	logger.Logger.Info("Build info", "date", buildDate)

	a, err := app.Make()
	if err != nil {
		logger.Logger.Error("Error when make app", "err", err)
		os.Exit(1)
	}

	a.Start()
}
