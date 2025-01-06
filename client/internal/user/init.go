package user

import (
	"client/internal/config"
	"client/internal/models"
	"client/internal/roamer"
	"context"
)

type CallFunc func(f models.User) error

type UserAccessor interface {
	UserGet(ctx context.Context, login string) (user models.User, err error)
	UserCreate(ctx context.Context, u models.User) error
}

type Item struct {
	storage  UserAccessor
	settings config.MainConfig
	roamer   roamer.Item
}

func Make(stor UserAccessor, settings config.MainConfig, roamer roamer.Item) Item {
	return Item{
		storage: stor, settings: settings, roamer: roamer,
	}
}
