package syncer

import (
	"client/internal/config"
	"client/internal/models"
	"client/internal/versioning"
	"context"
)

type StorageAccessor interface {
	SecretsUpsert(ctx context.Context, list []models.Secret) error
	SecretList(ctx context.Context, userID string, lastSynced string) ([]models.Secret, error)

	GetLastSyncEventVersion(ctx context.Context, kind string) (version string, err error)
	SaveSyncEvent(ctx context.Context, kind string, version string) error
}

type RoamerAccessor interface {
	SecretGet(version string) (list []models.Secret, err error)
	SecretSet(list []models.Secret) error
}

type verGetter interface{ Get() string }

type Item struct {
	settings *config.MainConfig
	storage  StorageAccessor
	roamer   RoamerAccessor

	stoper chan bool
	verGet verGetter
}

func Make(settings *config.MainConfig, storage StorageAccessor, roamer RoamerAccessor) Item {
	return Item{
		settings: settings,
		storage:  storage,
		roamer:   roamer,

		stoper: make(chan bool),

		verGet: &versioning.Version{},
	}
}
