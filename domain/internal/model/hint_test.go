package model

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cwd, _   = os.Getwd()
	testdata = path.Join(cwd, "testdata")
)

func Load(name string) *Board {
	jsonData, _ := os.ReadFile(path.Join(testdata, name))

	var board Board
	json.Unmarshal(jsonData, &board)
	return &board
}

func TestHint(t *testing.T) {

	t.Run("GetPlaceTileHint", func(t *testing.T) {
		t.Run("GetPlaceTileHint should return the right hint", func(t *testing.T) {
			board := Load("place-tile-hint")

			updatedBoard, hint := board.GetPlaceTileHint()

			assert.NotEqual(t, board, updatedBoard)
			assert.Equal(t, &PlaceTileHint{Direction: "RIGHT", Index: 1, Rotation: 90}, hint)
		})

		t.Run("GetPlaceTileHint should return nil if no move lead to the goal", func(t *testing.T) {
			board := Load("place-tile-no-hint")

			updatedBoard, hint := board.GetPlaceTileHint()

			assert.Equal(t, board, updatedBoard)
			assert.Nil(t, hint)
		})
	})
}
