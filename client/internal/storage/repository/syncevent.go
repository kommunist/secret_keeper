package repository

import (
	"client/internal/logger"
	"context"
	"database/sql"
)

// TODO: в перспективе добавить еще и пользователя в эвент.
const getLastSyncEventQuery = `
  SELECT version from sync_events where kind = $1 order by version desc limit 1
`
const saveSyncEvent = `INSERT INTO sync_events (kind, version) VALUES ($1, $2)`

func (si *Storage) GetLastSyncEventVersion(ctx context.Context, kind string) (version string, err error) {
	row := si.driver.QueryRow(getLastSyncEventQuery, kind)
	err = row.Scan(&version)

	if err == nil {
		return version, nil
	}

	if err == sql.ErrNoRows {
		return "0", nil
	} else {
		return "", err
	}
}

func (si *Storage) SaveSyncEvent(ctx context.Context, kind string, version string) error {
	_, err := si.driver.ExecContext(ctx, saveSyncEvent, kind, version)
	if err != nil {
		logger.Logger.Error("error when save sync event", "err", err)
		return err
	}
	return nil
}
