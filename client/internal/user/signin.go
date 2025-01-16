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
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Logger.Info("try to athentificate user by server")
			u, err := i.roamer.UserGet(f)
			if err != nil {
				return err
			}

			err = i.storage.UserCreate(context.Background(), u)
			if err != nil {
				logger.Logger.Error("when create user locally", "err", err)
				return err
			}
			u.Password = f.Password

			current.SetUser(u)
			return nil
		}
		return err
	}

	logger.Logger.Info("user athentificated locally")
	u.Password = f.Password

	current.SetUser(u)
	return nil
}
