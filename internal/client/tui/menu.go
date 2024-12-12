package tui

import "github.com/rivo/tview"

func (t *Tui) Menu() {
	menu := []menuItem{
		{name: "Мои пароли", shortcut: 'a', target: nil},
		{name: "Мои документы", shortcut: 'b', target: nil},
		{name: "Выйти", shortcut: 'c', target: func() { t.application.Stop() }},
	}

	list := tview.NewList()
	for _, item := range menu {
		list.AddItem(item.name, "", item.shortcut, item.target)
	}
	t.Show(list)
}
