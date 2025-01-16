package encrypt

import "golang.org/x/crypto/bcrypt"

type Item struct{}

// Метод хеширования пароля
func (i Item) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Проверка, что пароль и хеш соответствуют
func (i Item) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
