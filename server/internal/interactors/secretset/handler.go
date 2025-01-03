package secretset

import (
	"io"
	"log/slog"
	"net/http"
	"server/internal/auth"
)

// Получает секреты от клиентов и сохраняет их в базу

func (h *Interactor) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value(auth.UserIDKey) == nil {
		return
	}
	userID := r.Context().Value(auth.UserIDKey).(string)

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("read body error", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = h.Perform(r.Context(), userID, body)
	if err != nil {
		slog.Error("error when perform request", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
