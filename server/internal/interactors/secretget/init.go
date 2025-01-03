package secretget

import (
	"context"
	"server/internal/config"
	"server/internal/models"
)

// Интерфейс доступа к базе
type SecretGetter interface {
	SecretGet(ctx context.Context, userID string, version string) ([]models.Secret, error)
}

// Структура хендлера
type Interactor struct {
	storage SecretGetter
	setting config.MainConfig
}

// Конструктор хендлера
func Make(setting config.MainConfig, storage SecretGetter) Interactor {
	return Interactor{setting: setting, storage: storage}
}
