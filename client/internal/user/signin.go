package user

import (
	"client/internal/logger"
	"client/internal/models"
	"context"
	"database/sql"
)

func (i *Item) SignIN(login string, password string) (models.User, error) {

	u, err := i.storage.UserGet(context.Background(), login)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Logger.Info("try to athentificate user by server")
			u, err := i.roamer.UserGet(models.User{Login: login, Password: password})
			if err != nil {
				return models.User{}, err
			}

			err = i.storage.UserCreate(context.Background(), u)
			if err != nil {
				logger.Logger.Error("when create user locally", "err", err)
				return models.User{}, err
			}
			u.Password = password

			return u, nil
		}
		return models.User{}, err
	}

	logger.Logger.Info("user athentificated locally")
	u.Password = password

	return u, nil
}
