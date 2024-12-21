package tui

import (
	"secret_keeper/internal/client/secret"

	"github.com/rivo/tview"
)

func (t *Tui) Menu() {
	menu := []menuItem{
		{name: "My secrets", shortcut: 'a', target: func() { t.SecretList() }},
		{name: "Add secret", shortcut: 'b', target: func() { t.SecretAddPage(secret.MakeForm()) }},
		{name: "Exit", shortcut: 'c', target: func() { t.application.Stop() }},
	}

	list := tview.NewList()
	for _, item := range menu {
		list.AddItem(item.name, "", item.shortcut, item.target)
	}
	t.Show(list)
}
