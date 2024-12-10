package tui

import (
	"github.com/rivo/tview"
)

type menuItem struct {
	name     string
	shortcut rune
	target   func()
}

func (t *Tui) Hello() {

	menu := []menuItem{
		{name: "Log IN", shortcut: 'a', target: t.LogIN},
		{name: "Sign UP", shortcut: 'b', target: t.SignUP},
	}

	list := tview.NewList()
	for _, item := range menu {
		if item.target != nil {
			list.AddItem(item.name, "", item.shortcut, item.target)
		}

	}
	t.Show(list)
}
