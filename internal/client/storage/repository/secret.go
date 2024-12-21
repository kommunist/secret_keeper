package repository

import (
	"context"
	"log/slog"
	"secret_keeper/internal/client/logger"
	"secret_keeper/internal/client/models"
)

func (si *Storage) SecretCreate(
	ctx context.Context,
	name string, pass string, meta string, userID string, version string,
) error {

	// TODO сделать обработку ситуации, когда такой name уже существует

	_, err := si.driver.ExecContext(
		ctx,
		"INSERT INTO secrets (name, pass, meta, user_id, version) VALUES ($1, $2, $3, $4, $5)",
		name, pass, meta, userID, version,
	)
	if err != nil {
		logger.Logger.Error("Error when insert secret", "err", err)
		return err
	}
	return nil
}

func (si *Storage) SecretList(ctx context.Context, userID string) ([]models.Secret, error) {
	rows, err := si.driver.QueryContext(
		ctx,
		"SELECT name, pass, meta from secrets where user_id = $1",
		userID,
	)

	if err != nil {
		logger.Logger.Error("Error when select secrets", "err", err)
		return []models.Secret{}, err
	}
	defer rows.Close()

	result := make([]models.Secret, 0)

	for rows.Next() {
		inst := models.MakeSecret()

		errScan := rows.Scan(&inst.Name, &inst.Pass, &inst.Meta)
		if errScan != nil {
			slog.Error("When scan data from select", "err", errScan)
			return nil, errScan
		}
		result = append(result, inst)
	}

	return result, nil
}
