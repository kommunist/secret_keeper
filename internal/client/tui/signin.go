package tui

import (
	"secret_keeper/internal/client/signin"

	"github.com/rivo/tview"
)

func (t *Tui) SignInForm(sif signin.Form) tview.Primitive {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("SignIN").SetTitleAlign(tview.AlignLeft)
	form = form.AddInputField("Login", "", 20, nil, func(text string) { sif.Login = text })
	form = form.AddInputField("Password", "", 20, nil, func(text string) { sif.Password = text })
	form = form.AddButton("Save", func() { t.SignInSaveButton(sif) })
	form = form.AddButton("Back", func() { t.Hello() })

	return form
}

func (t *Tui) SignInPage(sif signin.Form) {
	prim := t.SignInForm(sif)

	t.Show(prim)
}

func (t *Tui) SignInSaveButton(sif signin.Form) {
	err := t.signinFunc(sif)
	if err != nil {
		t.ErrorModal(err.Error(), t.SignInForm(sif))
	}
}
