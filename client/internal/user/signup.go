package user

import (
	"client/internal/encrypt"
	"client/internal/logger"
	"client/internal/models"
	"context"
)

func (i *Item) SignUP(u models.User) error {

	// TODO вкрутить логику взаимодействия с внешним сервером
	hashedPass, err := encrypt.HashPassword(u.Password)
	if err != nil {
		logger.Logger.Error("Error when hash password", "err", err)
		return err
	}

	err = i.storage.UserCreate(context.Background(), u.Login, hashedPass)
	if err != nil {
		logger.Logger.Error("Error when create user", "err", err)
		return err
	}

	return nil
}
