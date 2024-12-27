package user

import (
	"client/internal/encrypt"
	"client/internal/logger"
	"context"
)

func (i *Item) SignUP(f Form) error {
	// TODO вкрутить логику взаимодействия с внешним сервером
	hashedPass, err := encrypt.HashPassword(f.Password)
	if err != nil {
		logger.Logger.Error("Error when hash password", "err", err)
		return err
	}

	err = i.storage.UserCreate(context.Background(), f.Login, hashedPass)
	if err != nil {
		logger.Logger.Error("Error when create user", "err", err)
		return err
	}

	return nil
}
