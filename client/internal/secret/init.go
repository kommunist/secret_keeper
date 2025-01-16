package secret

import (
	"client/internal/models"
	"client/internal/versioning"
	"context"
)

type SecretAccessor interface {
	SecretsUpsert(ctx context.Context, list []models.Secret) error
	SecretList(ctx context.Context, userID string, lastSynced string) ([]models.Secret, error)
	SecretShow(ctx context.Context, ID string) (models.Secret, error)
}
type VerGetter interface{ Get() string }

type Item struct {
	storage SecretAccessor

	verGet VerGetter
}

func Make(stor SecretAccessor) Item {
	return Item{
		storage: stor,
		verGet:  &versioning.Version{},
	}
}
