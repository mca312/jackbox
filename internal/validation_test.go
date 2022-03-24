package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidEmail(t *testing.T) {
	t.Run("email with one @ is valid", func(t *testing.T) {
		r := IsValidEmail("valid@email")
		assert.True(t, r)
	})

	t.Run("email with two @ is invalid", func(t *testing.T) {
		r := IsValidEmail("invalid@email@toomany")
		assert.False(t, r)
	})

	t.Run("email with no @ is invalid", func(t *testing.T) {
		r := IsValidEmail("noATemail")
		assert.False(t, r)
	})

	t.Run("email with no length is invalid", func(t *testing.T) {
		r := IsValidEmail("")
		assert.False(t, r)
	})
}
