package secret

import (
	"context"
	"fmt"
	"secret_keeper/internal/client/current"
	"secret_keeper/internal/client/logger"
	"time"
)

func (i *Item) Create(f Form) error {
	// TODO реализовать обработку ошибки, когда логин уже существует
	err := i.storage.SecretCreate(
		context.Background(), f.Name, f.Pass, f.Meta, current.User.ID, fmt.Sprintf("%v", time.Now().Unix()),
	)
	if err != nil {
		logger.Logger.Error("Error when create secret", "err", err)
		return err
	}

	return nil
}
