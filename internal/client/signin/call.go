package signin

import (
	"context"
	"errors"
	"secret_keeper/internal/client/current"
	"secret_keeper/internal/client/encrypt"
	"secret_keeper/internal/client/errors/incorrectpass"
	"secret_keeper/internal/client/logger"
)

func (i *Item) Call(f Form) error {
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
