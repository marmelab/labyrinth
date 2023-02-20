package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	t.Run("Should return an error if size is even.", func(t *testing.T) {
		board, err := NewBoard(2)
		assert.NotNil(t, err)
		assert.Equal(t, "size must be an odd number, got: 2", err.Error())
		assert.Nil(t, board)
	})

	t.Run("Should return a board instance if size is odd", func(t *testing.T) {
		board, err := NewBoard(3)
		assert.Nil(t, err)
		assert.NotNil(t, board)
	})
}
