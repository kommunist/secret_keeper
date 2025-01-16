package tui

import (
	"client/internal/models"

	"github.com/rivo/tview"
)

func (t *Tui) SignUPForm(suf models.User) tview.Primitive {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("SignUP").SetTitleAlign(tview.AlignLeft)
	form = form.AddInputField("Login", suf.Login, 20, nil, func(text string) { suf.Login = text })
	form = form.AddInputField("Password", suf.Password, 20, nil, func(text string) { suf.Password = text })

	form = form.AddButton("Save", func() { t.SignUPSaveButton(suf) })
	form = form.AddButton("Back", func() { t.Hello() })

	return form
}

func (t *Tui) SignUPPage(suf models.User) {
	prim := t.SignUPForm(suf)

	t.Show(prim)
}

func (t *Tui) SignUPSaveButton(suf models.User) {
	err := t.signupFunc(suf)
	if err != nil {
		t.ErrorModal(err.Error(), t.SignUPForm(suf))
	} else {
		t.Hello()
	}
}
