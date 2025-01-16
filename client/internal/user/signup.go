package user

import (
	"client/internal/logger"
	"client/internal/models"
)

// Регистрация пользователя. Всегда только он-лайн
func (i *Item) SignUP(u models.User) error {
	err := i.roamer.UserSet(u)
	if err != nil {
		logger.Logger.Error("error when call roamer", "err", err)
		return err
	}

	return nil
}
