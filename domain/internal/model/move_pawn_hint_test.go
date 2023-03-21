package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovePawnHint(t *testing.T) {

	t.Run("GetMovePawnHint", func(t *testing.T) {
		t.Run("GetMovePawnHint should return the target if in path", func(t *testing.T) {
			board := Load("place-tile-hint")

			hint, err := board.GetMovePawnHint()

			assert.Equal(t, ErrInvalidAction, err)
			assert.Nil(t, hint)
		})

		t.Run("GetMovePawnHint should return the target if in path", func(t *testing.T) {
			board := Load("move-pawn-hint-shortest")

			hint, _ := board.GetMovePawnHint()

			assert.Equal(t, &Coordinate{0, 2}, hint)
		})

		t.Run("GetMovePawnHint should return the closest tile to the target if not in path", func(t *testing.T) {
			board := Load("move-pawn-hint-closest")

			hint, _ := board.GetMovePawnHint()

			assert.Equal(t, &Coordinate{6, 3}, hint)
		})

		t.Run("GetMovePawnHint should return the closest tile to the center if remaining tile", func(t *testing.T) {
			board := Load("move-pawn-hint-remaining")

			hint, _ := board.GetMovePawnHint()

			assert.Equal(t, &Coordinate{3, 4}, hint)
		})
	})
}
