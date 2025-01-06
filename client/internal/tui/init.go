package tui

import (
	"client/internal/logger"
	"client/internal/secret"
	"client/internal/user"

	"github.com/rivo/tview"
)

type Tui struct {
	application *tview.Application

	box tview.Primitive

	signupFunc user.CallFunc
	signinFunc user.CallFunc

	// functions for secret
	createSecretFunc secret.CallFunc
	listSecretFunc   secret.ListFunc
	showSecretFunc   secret.ShowFunc
}

func Make(
	signupFunc user.CallFunc,
	signinFunc user.CallFunc,
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

func (t *Tui) Start() error {
	err := t.application.Run()
	if err != nil {
		logger.Logger.Error("Error when start tui", "err", err)
		return err
	}
	return nil
}

func (t *Tui) Stop() {
	t.application.Stop()
}

func (t *Tui) Show(item tview.Primitive) {
	t.application.SetRoot(item, true).SetFocus(item)
}
