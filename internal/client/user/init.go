package user

import (
	"context"
)

type CallFunc func(f Form) error

type UserAccessor interface {
	UserGet(ctx context.Context, login string) (userID string, hashedPass string, err error)
	UserCreate(ctx context.Context, login string, password string) error
}

type Item struct {
	storage UserAccessor
}

func Make(stor UserAccessor) Item {
	return Item{
		storage: stor,
	}
}
