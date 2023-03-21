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
		tests := []struct {
			name            string
			wasBoardUpdated bool
			hint            *PlaceTileHint
		}{
			{"place-tile-hint", true, &PlaceTileHint{Direction: "RIGHT", Index: 1, Rotation: 90}},
			{"place-tile-no-hint", false, nil},
		}

		for _, test := range tests {
			board := Load(test.name)

			updatedBoard, hint := board.GetPlaceTileHint()
			if test.wasBoardUpdated {
				assert.NotEqual(t, board, updatedBoard)
			} else {
				assert.Equal(t, board, updatedBoard)
			}
			assert.Equal(t, test.hint, hint)
		}
	})
}
