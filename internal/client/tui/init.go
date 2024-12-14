package tui

import (
	"secret_keeper/internal/client/logger"
	"secret_keeper/internal/client/signin"
	"secret_keeper/internal/client/signup"

	"github.com/rivo/tview"
)

type Tui struct {
	application *tview.Application

	box tview.Primitive
	// Pages tview.Pages

	signupFunc signup.CallFunc
	signinFunc signin.CallFunc
}

func Make(signupFunc signup.CallFunc, signinFunc signin.CallFunc) Tui {
	application := tview.NewApplication()
	box := tview.NewBox().SetBorder(true).SetTitle("secret_keeper_client")

	return Tui{
		box:         box,
		application: application,
		signupFunc:  signupFunc,
		signinFunc:  signinFunc,
	}
}

func (t *Tui) Start() {
	err := t.application.Run()
	if err != nil {
		logger.Logger.Error("Error when start tui", "err", err)
		panic(err) // вытащитьб обработку ошибки
	}
}

func (t *Tui) Show(item tview.Primitive) {
	t.application.SetRoot(item, true).SetFocus(item)
}
