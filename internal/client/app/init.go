package app

import (
	"secret_keeper/internal/client/logger"
	"secret_keeper/internal/client/secret"
	"secret_keeper/internal/client/signin"
	"secret_keeper/internal/client/signup"
	"secret_keeper/internal/client/storage/repository"
	"secret_keeper/internal/client/tui"
)

type Storager interface {
	signup.UserCreator
	signin.UserGetter
	secret.SecretAccessor
}

type App struct {
	tui    tui.Tui     // Морда в терминале
	signup signup.Item // Сервис регистрации
	signin signin.Item // Сервис регистрации

	storage Storager // База
}

func Make() (App, error) {

	// TODO переделать на нормальный конфиг
	stor, err := repository.Make("postgresql://postgres:postgres@localhost:5435/secret_keeper")
	if err != nil {
		logger.Logger.Error("Error when make storage", "err", err)
		return App{}, err
	}

	signUPItem := signup.Make(&stor)
	signINItem := signin.Make(&stor)
	secretItem := secret.Make(&stor)

	tuiItem := tui.Make(
		signUPItem.Call, // метод для регистрации
		signINItem.Call, // метод для логина
		secretItem.Create,
		secretItem.List,
		secretItem.Show,
	)

	return App{
		tui:     tuiItem,
		signup:  signUPItem,
		signin:  signINItem,
		storage: &stor,
	}, nil
}

func (a *App) Start() {
	a.tui.Hello()
	a.tui.Start()
}

func (a *App) Stop() {
	a.tui.Stop()
}
