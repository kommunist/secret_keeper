package app

import (
	"client/internal/config"
	"client/internal/logger"
	"client/internal/roamer"
	"client/internal/secret"
	"client/internal/storage/repository"
	"client/internal/syncer"
	"client/internal/tui"
	"client/internal/user"
)

type Storager interface {
	user.UserAccessor
	secret.SecretAccessor
	syncer.StorageAccessor
}

type App struct {
	tui tui.Tui // Морда в терминале

	storage Storager // База
	syncer  syncer.Item
}

func Make() (App, error) {
	conf := config.Make()
	conf.Init()

	stor, err := repository.Make(conf.DatabaseURI)

	if err != nil {
		logger.Logger.Error("Error when make storage", "err", err)
		return App{}, err
	}

	roamer := roamer.Make(conf)
	syncer := syncer.Make(conf, &stor, &roamer)
	userItem := user.Make(&stor, &roamer)
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
		syncer:  syncer,
	}, nil
}

func (a *App) Start() error {
	a.syncer.Start()

	a.tui.Hello()
	err := a.tui.Start()
	if err != nil {
		logger.Logger.Error("Start: when start tui", "err", err)
		return err
	}
	return nil
}

func (a *App) Stop() {
	a.tui.Stop()
}
