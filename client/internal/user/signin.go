package user

import (
	"client/internal/current"
	"client/internal/encrypt"
	"client/internal/errors/incorrectpass"
	"client/internal/logger"
	"client/internal/models"
	"context"
	"errors"
)

func (i *Item) SignIN(f models.User) error {
	userID, hashedPass, err := i.storage.UserGet(context.Background(), f.Login)
	if err != nil {
		logger.Logger.Error("Error when get user from storage", "err", err)
		return err
	}

	if !encrypt.CheckPasswordHash(f.Password, hashedPass) {
		return incorrectpass.NewIncorrectPassError(errors.New("incorrect pass"))
	}

	current.SetUser(f.Login, f.Password, userID)

	return nil
}
