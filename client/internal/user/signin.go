package user

import (
	"client/internal/current"
	"client/internal/logger"
	"client/internal/models"
	"context"
	"database/sql"
)

func (i *Item) SignIN(f models.User) error {

	u, err := i.storage.UserGet(context.Background(), f.Login)
	if err == nil {
		logger.Logger.Info("user athentificated locally")
		u.Password = f.Password

		current.SetUser(u)
		return nil
	}
	if err == sql.ErrNoRows {
		logger.Logger.Info("try to athentificate user by server")
		u, roerr := i.roamer.UserGet(f)
		if roerr == nil {
			i.storage.UserCreate(context.Background(), u)
			u.Password = f.Password

			logger.Logger.Info("Current user bore auth", "user", u)

			current.SetUser(u)
			return nil
		}
		return roerr
	}
	return err
}
