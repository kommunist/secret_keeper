package secretget

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"server/internal/auth"
)

// Вытаскивает секреты из базы для текущего пользователя
func (h *Interactor) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Context().Value(auth.UserIDKey) == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userID := r.Context().Value(auth.UserIDKey).(string)

	version := r.URL.Query().Get("version")

	list, err := h.storage.SecretGet(r.Context(), userID, version)
	if err != nil {
		slog.Error("error when get secrets from storage", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(list)
	if err != nil {
		slog.Error("error when convert result to json", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		slog.Error("error when write data to response", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
