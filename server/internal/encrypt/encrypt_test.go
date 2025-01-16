package encrypt

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndCheck(t *testing.T) {
	i := Item{}

	password := "password"
	encrypted, err := i.HashPassword(password)

	t.Run("it_correct_encrypt_and_check_password", func(t *testing.T) {
		assert.NoError(t, err, "hashed_password_without_error")

		assert.True(t, i.CheckPasswordHash(password, encrypted))

	})

	t.Run("it_correce_encrypt_and_check_password_when_password_incorrect", func(t *testing.T) {
		assert.NoError(t, err, "hashed_password_without_error")

		assert.False(t, i.CheckPasswordHash(strings.ToUpper(password), encrypted))
	})

}
