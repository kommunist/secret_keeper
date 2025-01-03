package app

import (
	"github.com/go-chi/chi/v5"
)

func (a *App) initRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Post("/users", a.userset.Handler) // Регистрация

		r.With(a.authCheck.Check).Group(func(r chi.Router) {
			r.Post("/secrets", a.secretset.Handler) // Загрузка секретов
			r.Get("/secrets", a.secretget.Handler)  // Получение секретов
			r.Post("/auth", a.userget.Handler)      // сознательно выбран POST, чтобы не тащить ничего в урле

		})
	})

	return r
}
