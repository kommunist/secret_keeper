package repository

import (
	"context"
	"secret_keeper/internal/client/logger"
)

func (si *Storage) UserCreate(ctx context.Context, login string, password string) error {

	// TODO сделать обработку ситуации, когда пользователь уже существует

	_, err := si.driver.ExecContext(
		ctx,
		"INSERT INTO users (login, password) VALUES ($1, $2)",
		login, password,
	)
	if err != nil {
		logger.Logger.Error("Error when insert user", "err", err)
		return err
	}
	return nil
}
