package signup

type Form struct {
	Login    string
	Password string
}

func MakeForm() Form {
	return Form{}
}
