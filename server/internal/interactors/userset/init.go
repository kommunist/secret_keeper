package userset

import (
	"context"
	"server/internal/config"
	"server/internal/models"
)

// Интерфейс доступа к базе
type UserSetter interface {
	UserSet(ctx context.Context, user models.User) error
}

// Структура хендлера
type Interactor struct {
	storage UserSetter
	setting config.MainConfig
}

// Конструктор хендлера
func Make(setting config.MainConfig, storage UserSetter) Interactor {
	return Interactor{setting: setting, storage: storage}
}
