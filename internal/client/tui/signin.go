package tui

import (
	"secret_keeper/internal/client/signin"

	"github.com/rivo/tview"
)

func (t *Tui) SignIN() {
	sif := signin.MakeForm()

	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	form = form.AddInputField("Login", "", 20, nil, func(text string) { sif.Login = text })
	form = form.AddInputField("Password", "", 20, nil, func(text string) { sif.Password = text })
	form = form.AddButton("Save", func() { t.SignINSaveButton(sif) })
	form = form.AddButton("Back", func() { t.Hello() })

	t.Show(form)
}

func (t *Tui) SignINSaveButton(sif signin.Form) {
	t.signinFunc(sif)
	t.Menu()
	// t.Hello()
}
