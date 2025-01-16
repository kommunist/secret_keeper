package secretset

import (
	"context"
	"encoding/json"
	"log/slog"
	"server/internal/models"
)

func (h *Interactor) Perform(ctx context.Context, userID string, body []byte) error {

	list := []models.Secret{}
	err := json.Unmarshal(body, &list)
	if err != nil {
		slog.Error("error when parse json", "err", err)
		return err
	}

	for ind := 0; ind < len(list); ind++ {
		list[ind].UserID = userID
	}

	err = h.storage.SecretUpsert(ctx, list)
	if err != nil {
		slog.Error("error when upsert secrets", "err", err)
		return err
	}

	return nil
}
