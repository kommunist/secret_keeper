package secret

import (
	"client/internal/logger"
	"client/internal/models"
	"context"
)

type ShowFunc func(ID string) (models.Secret, error)

func (i *Item) Show(ID string) (models.Secret, error) {
	item, err := i.storage.SecretShow(context.Background(), ID)
	if err != nil {
		logger.Logger.Error("Error when get list of secrets", "err", err)
		return models.Secret{}, err
	}

	return item, nil
}
