package tui

import (
	"secret_keeper/internal/client/logger"
	"secret_keeper/internal/client/secret"
	"secret_keeper/internal/client/signin"
	"secret_keeper/internal/client/signup"

	"github.com/rivo/tview"
)

// TODO придумать, как показывать ошибки в TUI. Сейчас пока сделал вывод ошибок только в
// логгер

type Tui struct {
	application *tview.Application

	box tview.Primitive

	signupFunc signup.CallFunc
	signinFunc signin.CallFunc

	// functions for secret
	createSecretFunc secret.CallFunc
	listSecretFunc   secret.ListFunc
	showSecretFunc   secret.ShowFunc
}

func Make(
	signupFunc signup.CallFunc,
	signinFunc signin.CallFunc,
	createSecretFunc secret.CallFunc,
	listSecretFunc secret.ListFunc,
	showSecretFunc secret.ShowFunc,
) Tui {
	application := tview.NewApplication()
	box := tview.NewBox().SetBorder(true).SetTitle("secret_keeper_client")

	return Tui{
		box:              box,
		application:      application,
		signupFunc:       signupFunc,
		signinFunc:       signinFunc,
		createSecretFunc: createSecretFunc,
		listSecretFunc:   listSecretFunc,
		showSecretFunc:   showSecretFunc,
	}
}

func (t *Tui) Start() {
	err := t.application.Run()
	if err != nil {
		logger.Logger.Error("Error when start tui", "err", err)
		panic(err) // TODO вытащить обработку ошибки
	}
}

func (t *Tui) Stop() {
	t.application.Stop()
}

func (t *Tui) Show(item tview.Primitive) {
	t.application.SetRoot(item, true).SetFocus(item)
}
