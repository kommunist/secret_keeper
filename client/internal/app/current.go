package app

import "client/internal/models"

func (a *App) setCurrent(u models.User) {
	a.current = u
}

func (a *App) getCurrent() models.User {
	return a.current
}
