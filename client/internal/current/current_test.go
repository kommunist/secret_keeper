package current

import (
	"client/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetUser(t *testing.T) {
	t.Run("is_set_current_user", func(t *testing.T) {
		defer UnsetUser()

		u := models.User{Login: "login", Password: "password", ID: "id"}

		SetUser(u)

		assert.Equal(t, u.Login, User.Login, "correct_set_login_of_user")
		assert.Equal(t, u.Password, User.Password, "correct_set_password_of_user")
		assert.Equal(t, u.ID, User.ID, "correct_set_id_of_user")

		assert.True(t, UserSeted(), "chech_by_special_method")
	})
}

func TestUnSetUser(t *testing.T) {
	t.Run("is_unset_current_user", func(t *testing.T) {
		defer UnsetUser()

		u := models.User{Login: "login", Password: "password", ID: "id"}
		SetUser(u)

		UnsetUser()

		assert.NotEqual(t, u.Login, User.Login, "correct_unset_login_of_user")
		assert.NotEqual(t, u.Password, User.Password, "correct_unset_password_of_user")
		assert.NotEqual(t, u.ID, User.ID, "correct_unset_id_of_user")
	})
}
