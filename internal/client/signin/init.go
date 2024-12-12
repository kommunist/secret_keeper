package signin

import (
	"context"
)

type CallFunc func(f Form) error

type UserGetter interface {
	UserGet(ctx context.Context, login string) (userID string, hashedPass string, err error)
}

type Item struct {
	storage UserGetter
}

func Make(stor UserGetter) Item {
	return Item{
		storage: stor,
	}
}
