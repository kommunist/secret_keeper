package tui

import (
	"secret_keeper/internal/client/user"

	"github.com/rivo/tview"
)

func (t *Tui) SignUPForm(suf user.Form) tview.Primitive {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("SignUP").SetTitleAlign(tview.AlignLeft)
	form = form.AddInputField("Login", suf.Login, 20, nil, func(text string) { suf.Login = text })
	form = form.AddInputField("Password", suf.Password, 20, nil, func(text string) { suf.Password = text })

	form = form.AddButton("Save", func() { t.SignUPSaveButton(suf) })
	form = form.AddButton("Back", func() { t.Hello() })

	return form
}

func (t *Tui) SignUPPage(suf user.Form) {
	prim := t.SignUPForm(suf)

	t.Show(prim)
}

func (t *Tui) SignUPSaveButton(suf user.Form) {
	err := t.signupFunc(suf)
	if err != nil {
		t.ErrorModal(err.Error(), t.SignUPForm(suf))
	} else {
		t.Hello()
	}
}
