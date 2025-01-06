package roamer

import (
	"client/internal/config"
	"net/http"
)

var client = &http.Client{}

type Item struct {
	settings config.MainConfig
}

func Make(settings config.MainConfig) Item {
	return Item{
		settings: settings,
	}
}
