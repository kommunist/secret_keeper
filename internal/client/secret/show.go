package secret

import (
	"context"
	"secret_keeper/internal/client/logger"
)

type ShowFunc func(ID string) (Form, error)

func (i *Item) Show(ID string) (Form, error) {
	item, err := i.storage.SecretShow(context.Background(), ID)
	if err != nil {
		logger.Logger.Error("Error when get list of secrets", "err", err)
		return Form{}, err
	}

	return Form(item), nil
}
