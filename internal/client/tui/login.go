package tui

import (
	"github.com/rivo/tview"
)

func (t *Tui) LogIN() {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)
	form = form.AddInputField("Login", "", 20, nil, nil)
	form = form.AddInputField("Password", "", 20, nil, nil)
	form = form.AddButton("Save", func() { t.application.Stop() })
	form = form.AddButton("Back", func() { t.Hello() })

	t.Show(form)
}
