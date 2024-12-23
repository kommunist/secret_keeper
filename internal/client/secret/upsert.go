package secret

import (
	"context"
	"fmt"
	"secret_keeper/internal/client/current"
	"secret_keeper/internal/client/logger"
	"time"
)

func (i *Item) Upsert(f Form) error {
	err := i.storage.SecretUpsert(
		context.Background(),
		f.ID, f.Name, f.Pass, f.Meta, current.User.ID, fmt.Sprintf("%v", time.Now().Unix()),
	)
	if err != nil {
		logger.Logger.Error("Error when create secret", "err", err)
		return err
	}

	return nil
}
