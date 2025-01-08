package userget

// Структура хендлера
type Interactor struct{}

// Конструктор хендлера
func Make() Interactor {
	return Interactor{}
}
