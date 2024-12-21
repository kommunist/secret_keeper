package models

type Secret struct {
	Name string
	Pass string
	Meta string
}

func MakeSecret() Secret {
	return Secret{}
}
