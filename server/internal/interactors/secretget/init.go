package secretget

import (
	"context"
	"server/internal/models"
)

// Интерфейс доступа к базе
type SecretGetter interface {
	SecretGet(ctx context.Context, userID string, version string) ([]models.Secret, error)
}

// Структура хендлера
type Interactor struct {
	storage SecretGetter
}

// Конструктор хендлера
func Make(storage SecretGetter) Interactor {
	return Interactor{storage: storage}
}
