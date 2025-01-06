package secret

import (
	"client/internal/current"
	"client/internal/logger"
	"client/internal/models"
	"context"

	"github.com/google/uuid"
)

func (i *Item) Upsert(f models.Secret) error {

	f.Version = i.verGet.Get()
	f.UserID = current.User.ID

	if f.ID == "" {
		f.ID = uuid.New().String()
	}

	err := i.storage.SecretsUpsert(context.Background(), []models.Secret{f})
	if err != nil {
		logger.Logger.Error("Error when create secret", "err", err)
		return err
	}

	return nil
}
