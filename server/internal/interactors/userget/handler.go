package userget

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h *Interactor) Handler(w http.ResponseWriter, r *http.Request) {
	// body, err := io.ReadAll(r.Body)
	// defer r.Body.Close()

	// TODO: получать получать current_user из аутентификации
	userID := r.URL.Query().Get("user_id")

	user, err := h.storage.UserGet(r.Context(), userID)

	if err != nil {
		slog.Error("error when get user from storage", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(user)
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
