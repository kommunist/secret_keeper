package secretset

import (
	"context"
	"server/internal/models"
)

// Интерфейс доступа к базе
type SecretUpserter interface {
	SecretUpsert(context.Context, []models.Secret) error
}

// Структура хендлера
type Interactor struct {
	storage SecretUpserter
}

// Конструктор хендлера
func Make(storage SecretUpserter) Interactor {
	return Interactor{storage: storage}
}
