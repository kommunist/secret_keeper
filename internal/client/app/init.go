package app

import (
	"secret_keeper/internal/client/signin"
	"secret_keeper/internal/client/signup"
	"secret_keeper/internal/client/storage/repository"
	"secret_keeper/internal/client/tui"

	"github.com/rivo/tview"
)

type App struct {
	tui    tui.Tui     // Морда в терминале
	signup signup.Item // Сервис регистрации
	signin signin.Item // Сервис регистрации

	storage repository.Storage // База
	pages   tview.Pages
}

func Make() App {

	// TODO переделать на нормальный конфиг
	stor := repository.Make("postgresql://postgres:postgres@localhost:5435/secret_keeper")
	signUPItem := signup.Make(&stor)
	signINItem := signin.Make(&stor)

	tuiItem := tui.Make(
		signUPItem.Call, // метод для регистрации
		signINItem.Call, // метод для логина
	)

	return App{
		tui:     tuiItem,
		signup:  signUPItem,
		signin:  signINItem,
		storage: stor,
	}
}

func (a *App) Start() {
	a.tui.Hello()
	a.tui.Start()
}
