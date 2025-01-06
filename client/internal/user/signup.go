package user

import (
	"client/internal/logger"
	"client/internal/models"
)

func (i *Item) SignUP(u models.User) error {
	err := i.roamer.UserSet(u)
	if err != nil {
		logger.Logger.Error("error when call roamer", "err", err)
		return err
	}

	return nil
}
