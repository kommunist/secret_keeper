package secret

import (
	"context"
	"secret_keeper/internal/client/models"
)

type CallFunc func(f Form) error

// Форма для создания и редактирования секрета. Добавлена для удобства работы с формой)
type Form struct {
	ID   string
	Name string
	Pass string
	Meta string
	Ver  string
}

func MakeForm() Form {
	return Form{}
}

func MakeFormFromModel(m models.Secret) Form {
	return Form(m)
}

type SecretAccessor interface {
	SecretCreate(ctx context.Context, name string, pass string, meta string, userID string, version string) error
	SecretList(ctx context.Context, userID string) ([]models.Secret, error)
	SecretShow(ctx context.Context, ID string) (models.Secret, error)
}

type Item struct {
	storage SecretAccessor
}

func Make(stor SecretAccessor) Item {
	return Item{
		storage: stor,
	}
}
