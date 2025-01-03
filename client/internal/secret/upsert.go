package secret

import (
	"client/internal/current"
	"client/internal/logger"
	"context"
	"fmt"
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
