package auth

import (
	"context"
	"server/internal/models"
)

// Используется для прокидывания id пользователя через контекст
type UserIDKeyType int

// Используется для прокидывания id пользователя через контекст
const UserIDKey UserIDKeyType = 0

type UserGetter interface {
	UserGet(ctx context.Context, login string) (user models.User, err error)
}

type Item struct {
	storage UserGetter
}

// Конструктор хендлера
func Make(storage UserGetter) Item {
	return Item{storage: storage}
}
