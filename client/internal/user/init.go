package user

import (
	"client/internal/models"
	"context"
)

type UserAccessor interface {
	UserGet(ctx context.Context, login string) (user models.User, err error)
	UserCreate(ctx context.Context, u models.User) error
}

type RemoteUserAccessor interface {
	UserSet(f models.User) error
	UserGet(f models.User) (user models.User, err error)
}

type currentSet func(models.User)

type Item struct {
	storage UserAccessor
	roamer  RemoteUserAccessor

	currentSetFunc currentSet
}

func Make(stor UserAccessor, roamer RemoteUserAccessor, currecurrentSetFunc currentSet) Item {
	return Item{
		storage: stor, roamer: roamer, currentSetFunc: currecurrentSetFunc,
	}
}
