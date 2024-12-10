package signup

import (
	"context"
	"secret_keeper/internal/client/logger"
)

func (i *Item) Call(f Form) error {
	// TODO подумать, откуда притащить контекст

	logger.Logger.Info("CReate method")

	err := i.storage.UserCreate(context.Background(), f.Login, f.Password)
	if err != nil {
		logger.Logger.Error("Error when create user", "err", err)
		return err
	}

	return nil
}
