package tui

import (
	"secret_keeper/internal/client/secret"

	"github.com/rivo/tview"
)

func (t *Tui) SecretAddForm(sf secret.Form) tview.Primitive {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Create new secret").SetTitleAlign(tview.AlignLeft)
	form = form.AddInputField("Login", sf.Name, 20, nil, func(text string) { sf.Name = text })
	form = form.AddInputField("Password", sf.Pass, 20, nil, func(text string) { sf.Pass = text })
	form = form.AddTextArea("Meta", sf.Meta, 20, 20, 100, func(text string) { sf.Meta = text })
	form = form.AddButton("Save", func() { t.SaveSecretButton(sf) })
	form = form.AddButton("Back", func() { t.Menu() })

	return form
}

func (t *Tui) SecretAddPage(sf secret.Form) {
	prim := t.SecretAddForm(sf)

	t.Show(prim)
}

func (t *Tui) SaveSecretButton(sf secret.Form) {
	err := t.createSecretFunc(sf)
	if err != nil {
		t.ErrorModal(err.Error(), t.SecretAddForm(sf))
	} else {
		t.Menu()
	}
}
