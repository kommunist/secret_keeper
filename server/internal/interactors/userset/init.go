package userset

import (
	"context"
	"server/internal/encrypt"
	"server/internal/models"
)

// Интерфейс доступа к базе
type UserSetter interface {
	UserSet(ctx context.Context, user models.User) error
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
}

// Структура хендлера
type Interactor struct {
	storage UserSetter

	hasher PasswordHasher
}

// Конструктор хендлера
func Make(storage UserSetter) Interactor {
	return Interactor{storage: storage, hasher: encrypt.Item{}}
}
