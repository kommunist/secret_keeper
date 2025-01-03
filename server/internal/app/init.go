package app

import (
	"log/slog"
	"net/http"
	"server/internal/auth"
	"server/internal/config"
	"server/internal/interactors/secretget"
	"server/internal/interactors/secretset"
	"server/internal/interactors/userget"
	"server/internal/interactors/userset"
	"server/internal/storage/repository"
)

// Основная структура
type App struct {
	setting config.MainConfig
	storage Storager

	secretset secretset.Interactor
	secretget secretget.Interactor

	userset userset.Interactor
	userget userget.Interactor

	authCheck auth.Item

	Server http.Server
}

// Собирательный интерфейс хранилища
type Storager interface {
	secretset.SecretUpserter
	secretget.SecretGetter
	userset.UserSetter
	userget.UserGetter

	auth.UserGetter
}

// Конструктор структуры
func Make() (*App, error) {
	c := config.Make()
	c.Init()

	rep, err := repository.Make(c.DatabaseURI)
	if err != nil {
		slog.Error("Error when open connect to storage", "err", err)
		return &App{}, nil
	}

	app := App{
		setting:   c,
		storage:   &rep,
		secretset: secretset.Make(c, &rep),
		secretget: secretget.Make(c, &rep),

		userset: userset.Make(c, &rep),
		userget: userget.Make(c, &rep),

		authCheck: auth.Make(&rep),
	}
	app.Server = http.Server{Addr: c.ServerURL, Handler: app.initRouter()}

	return &app, nil
}
