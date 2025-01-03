package repository

import (
	"client/internal/logger"
	"client/internal/models"
	"context"
	"log/slog"

	"github.com/google/uuid"
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

const listSQL = "SELECT id, name, pass, meta, version from secrets where user_id = $1"
const showSQL = "SELECT id, name, pass, meta, version from secrets where ID = $1"

// Метод, создающий/обновляющий секрет в базе
func (si *Storage) SecretUpsert(
	ctx context.Context,
	id string,
	name string, pass string, meta string, userID string, version string,
) error {

	if id == "" { // Если запись новая
		id = uuid.New().String()
	}

	_, err := si.driver.ExecContext(
		ctx, upsertSQL,
		id, name, pass, meta, userID, version,
	)
	if err != nil {
		logger.Logger.Error("Error when insert secret", "err", err)
		return err
	}
	return nil
}

// Метод, достающий секреты текущего пользователя
func (si *Storage) SecretList(ctx context.Context, userID string) ([]models.Secret, error) {
	rows, err := si.driver.QueryContext(ctx, listSQL, userID)

	if err != nil {
		logger.Logger.Error("Error when select secrets", "err", err)
		return []models.Secret{}, err
	}
	defer rows.Close()

	result := make([]models.Secret, 0)

	for rows.Next() {
		inst := models.MakeSecret()

		errScan := rows.Scan(&inst.ID, &inst.Name, &inst.Pass, &inst.Meta, &inst.Ver)
		if errScan != nil {
			slog.Error("When scan data from select", "err", errScan)
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

	err := row.Scan(&m.ID, &m.Name, &m.Pass, &m.Meta, &m.Ver)

	if err != nil {
		logger.Logger.Error("Error when select secrets", "err", err)
		return models.Secret{}, err
	}
	return m, nil
}
