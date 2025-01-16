package tui

import (
	"client/internal/current"
	"client/internal/models"

	"github.com/rivo/tview"
)

type menuItem struct {
	name     string
	shortcut rune
	target   func()
}

func (t *Tui) Hello() {
	if current.UserSeted() {
		t.Menu()
	} else {

		menu := []menuItem{
			{name: "Log IN", shortcut: 'a', target: func() { t.SignInPage(models.User{}) }},
			{name: "Sign UP", shortcut: 'b', target: func() { t.SignUPPage(models.User{}) }},
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
