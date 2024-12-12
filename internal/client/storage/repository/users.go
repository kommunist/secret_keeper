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

func (si *Storage) UserGet(ctx context.Context, login string) (userID string, hashedPass string, err error) {
	row := si.driver.QueryRowContext(ctx, `SELECT id, password FROM users WHERE login ilike $1 limit 1`, login)
	err = row.Scan(&userID, &hashedPass)
	if err != nil {
		logger.Logger.Error("Error when scan data from storage", "err", err)
		return "", "", err
	}

	return userID, hashedPass, err
}
