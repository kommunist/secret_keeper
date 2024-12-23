package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	t.Run("happy_path_init_config", func(t *testing.T) {
		t.Setenv("SERVER_URI", "server_uri")
		t.Setenv("DATABASE_URI", "database_uri")

		c := Make()
		c.Init()

		assert.Equal(t, "database_uri", c.DatabaseURI)
		assert.Equal(t, "server_uri", c.ServerURI)

	})
}
