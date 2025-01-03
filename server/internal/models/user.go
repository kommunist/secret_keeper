package models

type User struct {
	ID         string `json:"id"`
	Login      string `json:"name"`
	Password   string `json:"pass"`
	HashedPass string `json:"hashed_pass"`
}

func MakeUser() User { return User{} }
