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
		authPass := r.Header.Get("Password")

		model, err := i.storage.UserGet(ctx, authLogin)
		if err != nil {
			slog.Error("error when get user from storage", "err", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		i := encrypt.Item{}

		if ok := i.CheckPasswordHash(authPass, model.HashedPassword); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, UserIDKey, model)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(authFn)
}
