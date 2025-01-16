package encrypt

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndCheck(t *testing.T) {

	password := "password"
	encrypted, err := HashPassword(password)
	assert.NoError(t, err, "hashed_password_without_error")

	t.Run("it_correct_encrypt_and_check_password", func(t *testing.T) {
		assert.True(t, CheckPasswordHash(password, encrypted))
	})

	t.Run("it_correce_encrypt_and_check_password_when_password_incorrect", func(t *testing.T) {
		assert.False(t, CheckPasswordHash(strings.ToUpper(password), encrypted))
	})

}
