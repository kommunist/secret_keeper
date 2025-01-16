package syncer

import (
	"client/internal/logger"
	"context"
)

func (i *Item) syncSecrets() {
	if ok := i.currentGetFunc().ID == ""; ok {
		return
	}

	lastSynced, err := i.storage.GetLastSyncEventVersion(context.Background(), "secret")
	if err != nil {
		logger.Logger.Error("syncSecrets: when get last version", "err", err)
		return
	}

	// сначала отправили локальные секреты
	err = i.sendLocalSecrets(lastSynced)
	if err != nil {
		logger.Logger.Error("syncSecrets: when send local secrets to server", "err", err)
		return
	}

	// теперь забрали секреты с сервера
	err = i.getServerSecrets(lastSynced)
	if err != nil {
		logger.Logger.Error("syncSecrets: when get server secrets", "err", err)
		return
	}

	// сохранили точку синхронизации
	err = i.saveSyncEvent()
	if err != nil {
		logger.Logger.Error("syncSecrets: when save sync event", "err", err)
		return
	}
}

func (i *Item) sendLocalSecrets(lastSynced string) error {
	secrets, err := i.storage.SecretList(context.Background(), i.currentGetFunc().ID, lastSynced)
	if err != nil {
		logger.Logger.Error("sendLocalSecrets: when get secrets from local storage", "err", err)
		return err
	}
	logger.Logger.Info("sendLocalSecrets: send local secrets", "num", len(secrets))
	err = i.roamer.SecretSet(secrets, i.currentGetFunc())
	if err != nil {
		logger.Logger.Error("sendLocalSecrets: from roamer when send data to storage")
		return err
	}
	return nil
}

func (i *Item) getServerSecrets(version string) error {
	secrets, err := i.roamer.SecretGet(version, i.currentGetFunc())
	if err != nil {
		logger.Logger.Error("getServerSecrets: when get secrets from roamer", "err", err)
		return err
	}
	logger.Logger.Info("getServerSecrets: received from server secrets", "num", len(secrets))
	err = i.storage.SecretsUpsert(context.Background(), secrets)
	if err != nil {
		logger.Logger.Error("getServerSecrets: when upsert secrets to storage", "err", err)
		return err
	}
	return nil
}

func (i *Item) saveSyncEvent() error {
	eventVer := i.verGet.Get()
	err := i.storage.SaveSyncEvent(context.Background(), "secret", eventVer)
	if err != nil {
		logger.Logger.Info("saveSyncEvent: synced data", "version", eventVer)
		return err
	}
	return nil

}
