package signup

import (
	"context"
)

type CallFunc func(f Form) error

type UserCreator interface {
	UserCreate(ctx context.Context, login string, password string) error
}

type Item struct {
	storage UserCreator
}

func Make(stor UserCreator) Item {
	return Item{
		storage: stor,
	}
}
