package secretset

import (
	"io"
	"log/slog"
	"net/http"
	"server/internal/auth"
	"server/internal/models"
)

// Получает секреты от клиентов и сохраняет их в базу
func (h *Interactor) Handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handle request of secretset")

	if r.Context().Value(auth.UserIDKey) == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	user := r.Context().Value(auth.UserIDKey).(models.User)

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("read body error", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = h.Perform(r.Context(), user.ID, body)
	if err != nil {
		slog.Error("error when perform request", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
