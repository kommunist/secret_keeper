package repository

import (
	"context"
	"log/slog"
	"server/internal/models"
)

func (si *Storage) UserSet(ctx context.Context, user models.User) error {
	_, err := si.driver.ExecContext(
		ctx,
		"INSERT INTO users (login, password) VALUES ($1, $2)",
		user.Login, user.HashedPass,
	)
	if err != nil {
		slog.Error("Error when insert user", "err", err)
		return err
	}
	return nil
}

// TODO: подумать, действительно ли тут нужна целая модель для возврата
func (si *Storage) UserGet(ctx context.Context, login string) (user models.User, err error) {
	user = models.MakeUser()
	row := si.driver.QueryRowContext(ctx, `SELECT id, password FROM users WHERE login ilike $1 limit 1`, login)
	err = row.Scan(&user.ID, &user.HashedPass)
	if err != nil {
		slog.Error("Error when scan data from storage", "err", err)
		return models.User{}, err
	}

	return user, err
}
