package secret

import (
	"client/internal/models"
	"context"
)

type CallFunc func(s models.Secret) error

type SecretAccessor interface {
	SecretsUpsert(ctx context.Context, list []models.Secret) error
	SecretList(ctx context.Context, userID string, lastSynced string) ([]models.Secret, error)
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
