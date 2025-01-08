package userset

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"server/internal/models"
)

func (h *Interactor) Handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handle request of userset")

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("read body error", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		slog.Error("error when parse json body", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// если логика обработки вырастет, то целесообразно вытащить в отдельный сервис
	user.HashedPassword, err = h.hasher.HashPassword(user.Password)
	if err != nil {
		slog.Error("error when hash password", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = h.storage.UserSet(r.Context(), user)
	if err != nil {
		slog.Error("error when perform secrets", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
