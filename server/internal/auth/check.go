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

		authLogin := r.Header.Get("Login")
		authPass := r.Header.Get("Pass")

		model, err := i.storage.UserGet(ctx, authLogin)
		if err != nil {
			slog.Error("error when get user from storage", "err", err)
			return
		}

		if encrypt.CheckPasswordHash(authPass, model.HashedPass) {
			ctx = context.WithValue(ctx, UserIDKey, model.ID)

		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(authFn)
}
