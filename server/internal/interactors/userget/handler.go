package userget

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"server/internal/auth"
	"server/internal/models"
)

func (h *Interactor) Handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handle request of userget")

	if r.Context().Value(auth.UserIDKey) == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	user := r.Context().Value(auth.UserIDKey).(models.User)

	// user, err := h.storage.UserGet(r.Context(), userID)

	// if err != nil {
	// 	slog.Error("error when get user from storage", "err", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	body, err := json.Marshal(user)
	if err != nil {
		slog.Error("error when convert result to json", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		slog.Error("error when write data to response", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
