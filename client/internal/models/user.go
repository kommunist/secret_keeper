package models

// TODO при регистрации мы должны отправлять только login и password
type User struct {
	ID             string `json:"id"`
	Login          string `json:"login"`
	Password       string `json:"password"`
	HashedPassword string `json:"hashed_password"`
}

func MakeUser() User {
	return User{}
}
