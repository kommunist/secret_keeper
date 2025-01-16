package secret

import (
	"client/internal/logger"
	"client/internal/models"
	"context"
)

func (i *Item) List(u models.User) ([]models.Secret, error) {
	list, err := i.storage.SecretList(
		context.Background(), u.ID, "0",
	)
	if err != nil {
		logger.Logger.Error("Error when get list of secrets", "err", err)
		return nil, err
	}

	return list, nil
}
