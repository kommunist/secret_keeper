package roamer

import (
	"client/internal/logger"
	"client/internal/models"
	"encoding/json"
	"io"
	"net/http"
)

func (i *Item) SecretGet(version string, u models.User) (list []models.Secret, err error) {
	req, err := http.NewRequest("GET", i.settings.ServerURL+"/api/secrets", nil)
	if err != nil {
		logger.Logger.Error("error when prepare get secrets request to server", "err", err)
		return []models.Secret{}, err
	}

	q := req.URL.Query()
	q.Add("version", version)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Login", u.Login)
	req.Header.Set("Password", u.Password)

	resp, err := client.Do(req)
	if err != nil {
		logger.Logger.Error("error when send get secrets request to server", "err", err)
		return []models.Secret{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	logger.Logger.Info("response body from server", "body", string(body))

	err = json.Unmarshal(body, &list)
	if err != nil {
		logger.Logger.Error("error when parse json", "err", err)
		return []models.Secret{}, err
	}

	return list, nil
}
