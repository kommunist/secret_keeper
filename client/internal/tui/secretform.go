package tui

import (
	"client/internal/models"

	"github.com/rivo/tview"
)

// Форма для работы с секретом
func (t *Tui) SecretForm(sf models.Secret) tview.Primitive {
	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Create new secret").SetTitleAlign(tview.AlignLeft)
	form = form.AddTextView("ID", sf.ID, 40, 1, false, false)
	form = form.AddTextView("ver", sf.Ver, 40, 1, false, false)

	form = form.AddInputField("Login", sf.Name, 20, nil, func(text string) { sf.Name = text })
	form = form.AddInputField("Password", sf.Pass, 20, nil, func(text string) { sf.Pass = text })
	form = form.AddTextArea("Meta", sf.Meta, 20, 20, 100, func(text string) { sf.Meta = text })
	form = form.AddButton("Save", func() { t.SecretSaveButton(sf) })
	form = form.AddButton("Back", func() { t.Menu() })

	return form
}

// Страница работы с новым секретом
func (t *Tui) SecretCreatePage(sf models.Secret) {
	prim := t.SecretForm(sf)

	t.Show(prim)
}

func (t *Tui) SecretUpdatePage(ID string) {
	sf, err := t.showSecretFunc(ID)
	if err != nil {
		// TODO: Обработать ошибку модалкой
		panic(err)
	} else {
		prim := t.SecretForm(sf)

		t.Show(prim)
	}
}

// Кнопка сохранения на форме
func (t *Tui) SecretSaveButton(sf models.Secret) {
	err := t.createSecretFunc(sf)
	if err != nil {
		t.ErrorModal(err.Error(), t.SecretForm(sf))
	} else {
		t.Menu()
	}
}
