package tui

import (
	"secret_keeper/internal/client/current"
	"secret_keeper/internal/client/logger"

	"github.com/rivo/tview"
)

type menuItem struct {
	name     string
	shortcut rune
	target   func()
}

func (t *Tui) Hello() {
	logger.Logger.Info("User setedd", "setedd", current.UserSeted())
	if current.UserSeted() {
		t.Menu()
	} else {

		menu := []menuItem{
			{name: "Log IN", shortcut: 'a', target: t.SignIN},
			{name: "Sign UP", shortcut: 'b', target: t.SignUP},
			{name: "Exit", shortcut: 'e', target: func() { t.application.Stop() }},
		}

		list := tview.NewList()
		for _, item := range menu {
			if item.target != nil {
				list.AddItem(item.name, "", item.shortcut, item.target)
			}

		}
		t.Show(list)
	}
}
