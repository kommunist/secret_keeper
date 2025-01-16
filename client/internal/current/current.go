package current

// import (
// 	"client/internal/logger"
// 	"client/internal/models"
// )

// type UserInstance struct {
// 	ID       string
// 	Login    string
// 	Password string
// }

// // Здесь нарочно сделано как публичная глобальная переменная.
// // Предполагаю использовать из разных мест, где нужен текущий пользователь
// var User = UserInstance{}

// func SetUser(user models.User) {
// 	User = UserInstance{
// 		Login:    user.Login,
// 		Password: user.Password,
// 		ID:       user.ID,
// 	}
// 	logger.Logger.Info("User sign IN success", "ID", User.ID)
// }

// func UserSeted() bool {
// 	return User.ID != ""
// }

// func UnsetUser() {
// 	User = UserInstance{}
// }
