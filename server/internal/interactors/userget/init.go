package userget

import (
	"context"
	"server/internal/config"
	"server/internal/models"
)

// Интерфейс доступа к базе
type UserGetter interface {
	UserGet(ctx context.Context, userID string) (models.User, error)
}

// Структура хендлера
type Interactor struct {
	storage UserGetter
	setting config.MainConfig
}

// Конструктор хендлера
func Make(setting config.MainConfig, storage UserGetter) Interactor {
	return Interactor{setting: setting, storage: storage}
}
