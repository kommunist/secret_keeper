package user

import (
	"client/internal/logger"
	"client/internal/models"
	"context"
	"database/sql"
)

func (i *Item) SignIN(login string, password string) error {

	u, err := i.storage.UserGet(context.Background(), login)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Logger.Info("try to athentificate user by server")
			u, err := i.roamer.UserGet(models.User{Login: login, Password: password})
			if err != nil {
				return err
			}

			err = i.storage.UserCreate(context.Background(), u)
			if err != nil {
				logger.Logger.Error("when create user locally", "err", err)
				return err
			}
			u.Password = password

			i.currentSetFunc(u)

			return nil
		}
		return err
	}

	logger.Logger.Info("user athentificated locally")
	u.Password = password
	i.currentSetFunc(u)

	return nil
}
