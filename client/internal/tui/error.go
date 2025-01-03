package tui

import "github.com/rivo/tview"

func (t *Tui) ErrorModal(text string, prim tview.Primitive) {
	modal :=
		tview.NewModal().SetText(text).AddButtons([]string{"Back"}).
			SetDoneFunc(func(_ int, _ string) {
				t.Show(prim)
			})
	t.Show(modal)
}
