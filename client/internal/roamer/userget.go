package roamer

import (
	"client/internal/logger"
	"client/internal/models"
	"encoding/json"
	"io"
	"net/http"
)

// Метод для похода на сервер и получения данных о пользователе
func (i *Item) UserGet(f models.User) (user models.User, err error) {
	req, err := http.NewRequest("GET", i.settings.ServerURL+"/api/users", nil)
	if err != nil {
		logger.Logger.Error("error when prepare get user request to server", "err", err)
		return models.User{}, err
	}
	req.SetBasicAuth(f.Login, f.Password)

	resp, err := client.Do(req)
	if err != nil {
		logger.Logger.Error("error when send get user request to server", "err", err)
		return models.User{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	logger.Logger.Info("response body from server", "body", string(body))

	err = json.Unmarshal(body, &user)
	if err != nil {
		logger.Logger.Error("error when parse json", "err", err)
		return models.User{}, err
	}

	return user, nil
}
