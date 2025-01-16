package repository

import (
	"client/internal/logger"
	"client/internal/models"
	"context"
)

const userCreateSQL = "INSERT INTO users (id, login, password) VALUES ($1, $2, $3)"
const userGetSQL = "SELECT id, login, password FROM users WHERE login ilike $1 limit 1"

func (si *Storage) UserCreate(ctx context.Context, u models.User) error {
	_, err := si.driver.ExecContext(ctx, userCreateSQL, u.ID, u.Login, u.HashedPassword)
	if err != nil {
		logger.Logger.Error("Error when insert user", "err", err)
		return err
	}
	return nil
}

func (si *Storage) UserGet(ctx context.Context, login string) (user models.User, err error) {
	user = models.User{}

	err = si.driver.QueryRowContext(ctx, userGetSQL, login).Scan(&user.ID, &user.Login, &user.HashedPassword)
	if err != nil {
		logger.Logger.Error("Error when scan data from storage", "err", err)
		return models.User{}, err
	}

	return user, err
}
