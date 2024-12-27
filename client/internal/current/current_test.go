package current

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetUser(t *testing.T) {
	t.Run("is_set_current_user", func(t *testing.T) {
		defer UnsetUser()

		login := "login"
		password := "password"
		id := "id"

		SetUser(login, password, id)

		assert.Equal(t, login, User.Login, "correct_set_login_of_user")
		assert.Equal(t, password, User.Password, "correct_set_password_of_user")
		assert.Equal(t, id, User.ID, "correct_set_id_of_user")

		assert.True(t, UserSeted(), "chech_by_special_method")
	})
}

func TestUnSetUser(t *testing.T) {
	t.Run("is_unset_current_user", func(t *testing.T) {
		defer UnsetUser()

		login := "login"
		password := "password"
		id := "id"

		SetUser(login, password, id)
		UnsetUser()

		assert.NotEqual(t, login, User.Login, "correct_unset_login_of_user")
		assert.NotEqual(t, password, User.Password, "correct_unset_password_of_user")
		assert.NotEqual(t, id, User.ID, "correct_unset_id_of_user")
	})
}
