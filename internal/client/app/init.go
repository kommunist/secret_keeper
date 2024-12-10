package app

import (
	"secret_keeper/internal/client/signup"
	"secret_keeper/internal/client/storage/repository"
	"secret_keeper/internal/client/tui"
)

type App struct {
	tui     tui.Tui            // Морда в терминале
	signup  signup.Item        // Сервис регистрации
	storage repository.Storage // База
}

func Make() App {

	// TODO переделать на нормальный конфиг
	stor := repository.Make("postgresql://postgres:postgres@localhost:5435/secret_keeper")
	signUPItem := signup.Make(&stor)
	tuiItem := tui.Make(signUPItem.Call)

	return App{
		tui:     tuiItem,
		signup:  signUPItem,
		storage: stor,
	}
}

func (a *App) Start() {
	a.tui.Hello()
}
