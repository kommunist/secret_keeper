package secret

import (
	"client/internal/current"
	"client/internal/logger"
	"client/internal/models"
	"context"
)

type ListFunc func() ([]models.Secret, error)

func (i *Item) List() ([]models.Secret, error) {
	list, err := i.storage.SecretList(
		context.Background(), current.User.ID, "0",
	)
	if err != nil {
		logger.Logger.Error("Error when get list of secrets", "err", err)
		return nil, err
	}

	return list, nil
}
