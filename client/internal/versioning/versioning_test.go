package versioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Run("it_return_string_as_version", func(t *testing.T) {
		item := Version{}
		result := item.Get()
		assert.Len(t, result, 10) // 11ый разряд появится +- через 250 лет.
	})
}
