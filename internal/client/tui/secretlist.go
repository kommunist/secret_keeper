package tui

import (
	"fmt"
	"secret_keeper/internal/client/logger"
	"secret_keeper/internal/client/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Список секретов текущего пользователя
func (t *Tui) SecretList() {
	items, err := t.listSecretFunc()
	if err != nil {
		logger.Logger.Error("Error when get list of secrets", "err", err)
		t.Menu()
	}

	table :=
		tview.NewTable().
			SetBorders(true).
			SetSelectable(true, false).
			Select(0, 0).
			SetSelectedFunc(func(row int, column int) { t.secretListSelectedFunc(row, items) })
	buildHeader(table)

	for ind, item := range items {
		buildRow(ind, item, table)
	}
	buildCell(len(items)+1, 0, "Back", table)

	t.Show(table)
}

func buildHeader(table *tview.Table) {
	buildCell(0, 0, "index", table)
	buildCell(0, 1, "ID", table)
	buildCell(0, 2, "ver", table)
	buildCell(0, 3, "name", table)
	buildCell(0, 4, "pass", table)
	buildCell(0, 5, "meta", table)
}

func buildRow(ind int, item models.Secret, table *tview.Table) {
	buildCell(ind+1, 0, fmt.Sprintf("%d", ind), table)
	buildCell(ind+1, 1, item.ID, table)
	buildCell(ind+1, 2, item.Ver, table)
	buildCell(ind+1, 3, item.Name, table)
	buildCell(ind+1, 4, item.Pass, table)
	buildCell(ind+1, 5, item.Meta, table)
}

func buildCell(row int, column int, text string, table *tview.Table) {
	cell := tview.NewTableCell(text).SetTextColor(tcell.ColorGreenYellow).SetAlign(tview.AlignCenter)
	table.SetCell(row, column, cell)
}

func (t *Tui) secretListSelectedFunc(row int, items []models.Secret) {
	backRow := len(items) + 1

	switch row {
	case backRow:
		t.Menu()
	case 0:
	default:
		t.SecretUpdatePage(items[row-1].ID)
	}
}
