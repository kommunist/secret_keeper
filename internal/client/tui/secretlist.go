package tui

import (
	"fmt"
	"secret_keeper/internal/client/logger"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (t *Tui) SecretList() {
	items, err := t.listSecretFunc()
	if err != nil {
		logger.Logger.Error("Error when get list of secrets", "err", err)
		t.Menu()
	}

	table := tview.NewTable().SetBorders(true).SetSelectable(true, false).Select(0, 0)
	table.SetSelectedFunc(func(row int, column int) {
		t.secretListSelectedFunc(row, table, len(items)+1)
	})

	buildCell(0, 0, "index", table)
	buildCell(0, 1, "name", table)
	buildCell(0, 2, "pass", table)
	buildCell(0, 3, "meta", table)

	for ind, item := range items {
		buildCell(ind+1, 0, fmt.Sprintf("%d", ind), table)
		buildCell(ind+1, 1, item.Name, table)
		buildCell(ind+1, 2, item.Pass, table)
		buildCell(ind+1, 3, item.Meta, table)
	}
	buildCell(len(items)+1, 0, "Back", table)

	t.Show(table)
}

func buildCell(row int, column int, text string, table *tview.Table) {
	cell := tview.NewTableCell(text).SetTextColor(tcell.ColorGreenYellow).SetAlign(tview.AlignCenter)
	table.SetCell(row, column, cell)
}

func (t *Tui) secretListSelectedFunc(row int, table *tview.Table, backRow int) {
	if row == backRow {
		t.Menu()

	}
	// здесь нужно вариант для выбранной строки
}
