package tui

import (
	"secret_keeper/internal/client/logger"
	"secret_keeper/internal/client/signup"

	"github.com/rivo/tview"
)

type Tui struct {
	box         *tview.Box
	application *tview.Application

	signupFunc signup.CallFunc
}

func Make(signupFunc signup.CallFunc) Tui {
	box := tview.NewBox().SetBorder(true).SetTitle("Secret keeper client")
	application := tview.NewApplication()

	return Tui{
		box:         box,
		application: application,
		signupFunc:  signupFunc,
	}
}

func (t *Tui) Show(prim tview.Primitive) {

	err := t.application.SetRoot(prim, true).Run()
	if err != nil {
		logger.Logger.Error("Error when start tui", "err", err)
		panic(err) // вытащитьб обработку ошибки
	}
}
