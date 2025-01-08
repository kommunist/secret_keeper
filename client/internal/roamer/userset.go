package roamer

import (
	"bytes"
	"client/internal/logger"
	"client/internal/models"
	"encoding/json"
	"errors"
	"net/http"
)

func (i *Item) UserSet(f models.User) error {
	postBody, err := json.Marshal(f)

	if err != nil {
		logger.Logger.Error("Error when covert data to json", "err", err)
		return err
	}
	requestBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(i.settings.ServerURL+"/api/users", "application/json", requestBody)

	if err != nil {
		logger.Logger.Error("error when send request to server", "err", err)
		return err
	}

	// TODO в дальнейшем стоит развить систему статусов
	// (точно обработать ситуацию, когда пользователь существует)
	// пока оставляю в состоянии mvp
	if resp.StatusCode != http.StatusOK {
		logger.Logger.Info("server return not ok")
		return errors.New("server resurn not ok")
	}

	return nil
}
