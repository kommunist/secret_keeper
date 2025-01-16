package repository

import (
	"context"
	"log/slog"
	"server/internal/models"
)

const userSetSQL = "INSERT INTO users (login, password) VALUES ($1, $2)"
const userGetSQL = "SELECT id, login, password FROM users WHERE login ilike $1 limit 1"

func (si *Storage) UserSet(ctx context.Context, user models.User) error {
	_, err := si.driver.ExecContext(ctx, userSetSQL, user.Login, user.HashedPassword)
	if err != nil {
		slog.Error("error when insert user", "err", err)
		return err
	}
	return nil
}

func (si *Storage) UserGet(ctx context.Context, login string) (user models.User, err error) {
	user = models.User{}
	row := si.driver.QueryRowContext(ctx, userGetSQL, login)
	err = row.Scan(&user.ID, &user.Login, &user.HashedPassword)
	if err != nil {
		slog.Error("error when scan data from storage by user select", "err", err)
		return models.User{}, err
	}

	return user, err
}
