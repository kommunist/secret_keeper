package auth

import (
	"context"
	"log/slog"
	"net/http"
	"server/internal/encrypt"
)

// Мидлварь, который проверяет аутентификацию текущего пользователя
func (i *Item) Check(h http.Handler) http.Handler {
	authFn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		login, password, ok := r.BasicAuth()
		if !ok {
			slog.Info("Unauthorized request detected")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		model, err := i.storage.UserGet(ctx, login)
		if err != nil {
			slog.Error("error when get user from storage", "err", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		i := encrypt.Item{}

		if ok := i.CheckPasswordHash(password, model.HashedPassword); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, UserIDKey, model)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(authFn)
}
