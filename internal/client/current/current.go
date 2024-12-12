package current

import "secret_keeper/internal/client/logger"

type UserInstance struct {
	ID       string
	Login    string
	Password string
}

// Здесь нарочно сделано как публичная глобальная переменная.
// Предполагаю использовать из разных мест, где нужен текущий пользователь
var User = UserInstance{}

func SetUser(login string, password string, ID string) {
	User = UserInstance{
		Login:    login,
		Password: password,
		ID:       ID,
	}
	logger.Logger.Info("User sign IN success", "ID", User.ID)
}

func UserSeted() bool {
	return User.ID != ""
}
