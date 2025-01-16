package models

type User struct {
	ID             string `json:"id"`
	Login          string `json:"login"`
	Password       string `json:"password"`
	HashedPassword string `json:"hashed_password"`
}

type Secret struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Pass    string `json:"pass"`
	Meta    string `json:"meta"`
	UserID  string `json:"user_id"`
	Version string `json:"version"`
}
