package models

type Secret struct {
	ID   string
	Name string
	Pass string
	Meta string
	Ver  string
}

func MakeSecret() Secret {
	return Secret{}
}
