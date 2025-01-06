package syncer

import (
	"client/internal/config"
	"client/internal/logger"
	"client/internal/models"
	"client/internal/roamer"
	"context"
	"time"
)

type secretAccessor interface {
	SecretsUpsert(ctx context.Context, list []models.Secret) error
	SecretList(ctx context.Context, userID string, lastSynced string) ([]models.Secret, error)
}

type syncEventAccessor interface {
	GetLastSyncEventVersion(ctx context.Context, kind string) (version string, err error)
	SaveSyncEvent(ctx context.Context, kind string, version string) error
}

type SyncerStorageAccessor interface {
	secretAccessor
	syncEventAccessor
}

type Item struct {
	settings config.MainConfig
	storage  SyncerStorageAccessor
	roamer   roamer.Item

	stoper chan bool
}

func Make(settings config.MainConfig, storage SyncerStorageAccessor, roamer roamer.Item) Item {
	return Item{
		settings: settings,
		storage:  storage,
		roamer:   roamer,

		stoper: make(chan bool),
	}
}

func (i *Item) Start() {
	logger.Logger.Info("Syncer started")

	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				logger.Logger.Info("Syncer tick!")
				i.syncSecrets()
			case <-i.stoper:
				ticker.Stop()
				return
			}
		}
	}()
}
