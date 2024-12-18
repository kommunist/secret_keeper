package signup

import (
	"context"
	"secret_keeper/internal/client/encrypt"
	"secret_keeper/internal/client/logger"
)

func (i *Item) Call(f Form) error {
	hashedPass, err := encrypt.HashPassword(f.Password)
	if err != nil {
		logger.Logger.Error("Error when hash password", "err", err)
		return err
	}

	// TODO реализовать обработку ошибки, когда пользователь существует
	err = i.storage.UserCreate(context.Background(), f.Login, hashedPass)
	if err != nil {
		logger.Logger.Error("Error when create user", "err", err)
		return err
	}

	return nil
}
