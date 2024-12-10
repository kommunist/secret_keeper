package tui

import (
	"secret_keeper/internal/client/logger"
	"secret_keeper/internal/client/signup"

	"github.com/rivo/tview"
)

func (t *Tui) SignUP() {
	suf := signup.MakeForm()

	form := tview.NewForm()
	form.SetBorder(true).SetTitle("SignUP. Enter your data").SetTitleAlign(tview.AlignLeft)
	form = form.AddInputField("Login", "", 20, nil, func(text string) { suf.Login = text })
	form = form.AddInputField("Password", "", 20, nil, func(text string) { suf.Password = text })

	form = form.AddButton("Save", func() { t.SignUPSave(suf) }) // TODO написать полноценную функцию(с переходом)
	form = form.AddButton("Back", func() { t.Hello() })

	t.Show(form)
}

func (t *Tui) SignUPSave(suf signup.Form) {
	logger.Logger.Info("Save button pressed")
	t.signupFunc(suf)
	logger.Logger.Info("Form saved")
	t.Hello()

}
