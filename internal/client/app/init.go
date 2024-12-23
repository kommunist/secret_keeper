package app

import (
	"secret_keeper/internal/client/config"
	"secret_keeper/internal/client/logger"
	"secret_keeper/internal/client/secret"
	"secret_keeper/internal/client/storage/repository"
	"secret_keeper/internal/client/tui"
	"secret_keeper/internal/client/user"
)

type Storager interface {
	user.UserAccessor
	secret.SecretAccessor
}

type App struct {
	tui tui.Tui // Морда в терминале

	storage Storager // База
}

func Make() (App, error) {
	conf := config.Make()
	conf.Init()

	stor, err := repository.Make(conf.DatabaseURI)

	if err != nil {
		logger.Logger.Error("Error when make storage", "err", err)
		return App{}, err
	}

	userItem := user.Make(&stor)
	secretItem := secret.Make(&stor)

	tuiItem := tui.Make(
		userItem.SignUP,   // метод для регистрации
		userItem.SignIN,   // метод для логина
		secretItem.Upsert, // метод для создания/обновления секрета
		secretItem.List,   // метод для получения списка секретов
		secretItem.Show,   // метод для получения одного секрета
	)

	return App{
		tui:     tuiItem,
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
