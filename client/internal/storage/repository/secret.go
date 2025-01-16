package repository

import (
	"client/internal/logger"
	"client/internal/models"
	"context"
	"log/slog"
)

const upsertSQL = `
		INSERT INTO secrets (id, name, pass, meta, user_id, version)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT ( id ) DO UPDATE SET
		name = excluded.name,
		pass = excluded.pass,
		meta = excluded.meta,
		version = excluded.version
`

const listSQL = `
	SELECT id, name, pass, meta, version 
	FROM secrets 
	WHERE user_id = $1 AND version > $2
`
const showSQL = "SELECT id, name, pass, meta, version from secrets where ID = $1"

// Метод, создающий/обновляющий секрет в базе
func (si *Storage) SecretsUpsert(ctx context.Context, list []models.Secret) error {
	tx, err := si.driver.BeginTx(ctx, nil)
	if err != nil {
		logger.Logger.Error("SecretsUpsert: when open transaction to upsert secrets", "err", err)
		return err
	}

	for _, secret := range list {
		_, err = tx.ExecContext(
			ctx, upsertSQL,
			secret.ID, secret.Name, secret.Pass, secret.Meta, secret.UserID, secret.Version,
		)
		if err != nil {
			logger.Logger.Info("SecretsUpsert: list", "secrets", secret)
			logger.Logger.Error("SecretsUpsert: when exec upsert query", "err", err)
			tx.Rollback() // TODO игнорирую error
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		logger.Logger.Error("SecretsUpsert: when commit transaction on upsert secrets")
		tx.Rollback() // TODO игнорирую error
		return err
	}

	return nil
}

// Метод, достающий секреты текущего пользователя
func (si *Storage) SecretList(ctx context.Context, userID string, lastSynced string) ([]models.Secret, error) {
	rows, err := si.driver.QueryContext(ctx, listSQL, userID, lastSynced)

	if err != nil {
		logger.Logger.Error("SecretList: when select secrets", "err", err)
		return []models.Secret{}, err
	}
	defer rows.Close()

	result := make([]models.Secret, 0)

	for rows.Next() {
		inst := models.Secret{}

		errScan := rows.Scan(&inst.ID, &inst.Name, &inst.Pass, &inst.Meta, &inst.Version)
		if errScan != nil {
			slog.Error("SecretList: when scan data from select", "err", errScan)
			return nil, errScan
		}
		result = append(result, inst)
	}

	return result, nil
}

// Метод, позволяющий смотреть один конкретный секрет
func (si *Storage) SecretShow(ctx context.Context, ID string) (models.Secret, error) {
	m := models.Secret{}

	row := si.driver.QueryRowContext(ctx, showSQL, ID)

	err := row.Scan(&m.ID, &m.Name, &m.Pass, &m.Meta, &m.Version)

	if err != nil {
		logger.Logger.Error("SecretShow: when select secrets", "err", err)
		return models.Secret{}, err
	}
	return m, nil
}
