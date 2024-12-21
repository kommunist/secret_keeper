package tui

import (
	"secret_keeper/internal/client/secret"

	"github.com/rivo/tview"
)

func (t *Tui) SecretAdd() {
	sf := secret.MakeForm()

	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Create new secret").SetTitleAlign(tview.AlignLeft)
	form = form.AddInputField("Login", "", 20, nil, func(text string) { sf.Name = text })
	form = form.AddInputField("Password", "", 20, nil, func(text string) { sf.Pass = text })
	form = form.AddTextArea("Meta", "", 20, 20, 100, func(text string) { sf.Meta = text })
	form = form.AddButton("Save", func() { t.SaveSecretButton(sf) })
	form = form.AddButton("Back", func() { t.Menu() })

	t.Show(form)
}

func (t *Tui) SaveSecretButton(sf secret.Form) {
	err := t.createSecretFunc(sf)
	if err != nil {
		t.SecretAdd() // TODO: хорошо бы прокидывать данные, чтобы не терялись при ошибке
	} else {
		t.Menu()
	}
}
