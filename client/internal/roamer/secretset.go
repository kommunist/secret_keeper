package roamer

import (
	"bytes"
	"client/internal/logger"
	"client/internal/models"
	"encoding/json"
	"errors"
	"net/http"
)

func (i *Item) SecretSet(list []models.Secret, u models.User) error {
	postBody, err := json.Marshal(list)
	if err != nil {
		logger.Logger.Error("error when generate json to send data", "err", err)
		return err
	}
	requestBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest("POST", i.settings.ServerURL+"/api/secrets", requestBody)
	if err != nil {
		logger.Logger.Error("error when prepare post secrets request to server", "err", err)
		return err
	}
	req.SetBasicAuth(u.Login, u.Password)

	resp, err := client.Do(req)
	if err != nil {
		logger.Logger.Error("error when send post secrets request to server", "err", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Logger.Error("status of send secrets to server not OK", "status", resp.StatusCode)
		return errors.New("status not ok")
	}

	return nil
}
