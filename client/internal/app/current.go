package app

import "client/internal/models"

// метод задания текущего пользователя
func (a *App) setCurrent(u models.User) {
	a.current = u
}

// метод получения текущего пользователя
func (a *App) getCurrent() models.User {
	return a.current
}
