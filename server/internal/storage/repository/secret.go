package repository

import (
	"context"
	"log/slog"
	"server/internal/models"
)

const upsertSQL = `
		INSERT INTO secrets (id, name, pass, meta, user_id, version)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT ( id ) DO UPDATE SET
		name = excluded.name,
		pass = excluded.pass,
		meta = excluded.meta,
		version = excluded.version
		WHERE excluded.version > secrets.version
`

const getSQL = "SELECT id, name, pass, meta, version, user_id from secrets where user_id = $1 and version > $2"

// Метод, создающий/обновляющий секрет в базе
func (si *Storage) SecretUpsert(ctx context.Context, list []models.Secret) error {
	tx, err := si.driver.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("error when open transaction", "err", err)
		return err
	}

	for _, item := range list {
		_, err = tx.ExecContext(
			ctx, upsertSQL, item.ID, item.Name, item.Pass, item.Meta, item.UserID, item.Version,
		)

		if err != nil {
			slog.Error("error when insert data", "err", err)
			tx.Rollback() // TODO: игнорирую err
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("error when commit transaction on secrets", "err", err)
		return err
	}

	return nil
}

// Метод, достающий секреты текущего пользователя, но больше конерктной версии
func (si *Storage) SecretGet(ctx context.Context, userID string, version string) ([]models.Secret, error) {
	rows, err := si.driver.QueryContext(ctx, getSQL, userID, version)

	if err != nil {
		slog.Error("error when select secrets", "err", err)
		return []models.Secret{}, err
	}
	defer rows.Close()

	result := make([]models.Secret, 0)

	for rows.Next() {
		inst := models.Secret{}

		errScan := rows.Scan(&inst.ID, &inst.Name, &inst.Pass, &inst.Meta, &inst.UserID, &inst.Version)
		if errScan != nil {
			slog.Error("when scan data from select by secrets", "err", errScan)
			return nil, errScan
		}
		result = append(result, inst)
	}

	return result, nil
}
