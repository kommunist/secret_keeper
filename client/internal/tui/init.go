package tui

import (
	"client/internal/logger"
	"client/internal/models"

	"github.com/rivo/tview"
)

type signUPFuncType func(f models.User) error
type signINFuncType func(login string, password string) error

type secretCreateFunc func(f models.Secret, u models.User) error
type secretListFunc func(u models.User) ([]models.Secret, error)
type secretShowFunc func(ID string) (models.Secret, error)

type currentGet func() models.User

type Tui struct {
	application *tview.Application

	box tview.Primitive

	signupFunc signUPFuncType
	signinFunc signINFuncType

	createSecretFunc secretCreateFunc
	listSecretFunc   secretListFunc
	showSecretFunc   secretShowFunc

	currentGetFunc currentGet
}

func Make(
	signupFunc signUPFuncType,
	signinFunc signINFuncType,
	createSecretFunc secretCreateFunc,
	listSecretFunc secretListFunc,
	showSecretFunc secretShowFunc,

	currentGetFunc currentGet,
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

		currentGetFunc: currentGetFunc,
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
