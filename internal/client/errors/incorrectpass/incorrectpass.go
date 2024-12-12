package incorrectpass

// Основная структура ошибки.
type IncorrectPassError struct {
	Err error
}

// Вывод ошибки
func (e *IncorrectPassError) Error() string {
	return e.Err.Error()
}

// Инициализация ошибки
func NewIncorrectPassError(err error) error {
	return &IncorrectPassError{Err: err}
}
