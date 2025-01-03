package secret

import (
	"client/internal/models"
	"context"
)

type CallFunc func(s models.Secret) error

type SecretAccessor interface {
	SecretUpsert(ctx context.Context, id string, name string, pass string, meta string, userID string, version string) error
	SecretList(ctx context.Context, userID string) ([]models.Secret, error)
	SecretShow(ctx context.Context, ID string) (models.Secret, error)
}

type Item struct {
	storage SecretAccessor
}

func Make(stor SecretAccessor) Item {
	return Item{
		storage: stor,
	}
}
