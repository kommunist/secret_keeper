package tui

import (
	"secret_keeper/internal/client/signup"

	"github.com/rivo/tview"
)

func (t *Tui) SignUP() {
	suf := signup.MakeForm()

	form := tview.NewForm()
	form.SetBorder(true).SetTitle("SignUP").SetTitleAlign(tview.AlignLeft)
	form = form.AddInputField("Login", "", 20, nil, func(text string) { suf.Login = text })
	form = form.AddInputField("Password", "", 20, nil, func(text string) { suf.Password = text })

	form = form.AddButton("Save", func() { t.SignUPSaveButton(suf) })
	form = form.AddButton("Back", func() { t.Hello() })

	t.Show(form)
}

func (t *Tui) SignUPSaveButton(suf signup.Form) {
	t.signupFunc(suf)
	t.Hello()
}
