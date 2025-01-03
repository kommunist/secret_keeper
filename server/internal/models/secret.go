package models

type Secret struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Pass    string `json:"pass"`
	Meta    string `json:"meta"`
	UserID  string `json:"user_id"`
	Version string `json:"version"`
}

func MakeSecret() Secret { return Secret{} }
