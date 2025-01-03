package models

type User struct {
	Login    string
	Password string
}

func MakeUser() User {
	return User{}
}
