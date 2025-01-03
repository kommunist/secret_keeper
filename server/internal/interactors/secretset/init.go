package secretset

import (
	"context"
	"server/internal/config"
	"server/internal/models"
)

// Интерфейс доступа к базе
type SecretUpserter interface {
	SecretUpsert(context.Context, []models.Secret) error
}

// Структура хендлера
type Interactor struct {
	storage SecretUpserter
	setting config.MainConfig
}

// Конструктор хендлера
func Make(setting config.MainConfig, storage SecretUpserter) Interactor {
	return Interactor{setting: setting, storage: storage}
}
