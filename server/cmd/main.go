package main

import (
	"log/slog"
	"os"
	"server/internal/app"
)

var buildVersion string = "N/A"
var buildDate string = "N/A"

func main() {
	slog.Info("======Started======")
	slog.Info("Build info", "version", buildVersion)
	slog.Info("Build info", "date", buildDate)

	a, err := app.Make()
	if err != nil {
		slog.Error("Error when make app", "err", err)
		os.Exit(1)
	}

	a.Start()
}
