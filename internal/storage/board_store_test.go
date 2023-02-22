package storage

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/marmelab/labyrinth/internal/model"
	"github.com/stretchr/testify/assert"
)

var (
	cwd, _   = os.Getwd()
	testRoot = path.Join(cwd, "testdata")
)

func TestFileBoardStore(t *testing.T) {
	store := &fileBoardStore{
		rootDirectory: testRoot,
	}

	t.Run("Load()", func(t *testing.T) {
		t.Run("Should return an error if the state with the given id does not exist", func(t *testing.T) {
			board, err := store.Load("unexisting-board")
			assert.Nil(t, board)
			assert.NotNil(t, err)
			assert.True(t, os.IsNotExist(err))
		})

		t.Run("Should return an error if json parsing fails", func(t *testing.T) {
			board, err := store.Load("invalid-board")
			assert.Nil(t, board)
			assert.NotNil(t, err)
			assert.Equal(t, ErrMarshalFailure, err)
		})

		t.Run("Should return the board if found", func(t *testing.T) {
			board, err := store.Load("test-board")
			assert.Nil(t, err)
			assert.NotNil(t, board)

			assert.Equal(t, 3, len(board.Tiles))
			for line := 0; line < 3; line++ {
				assert.Equal(t, 3, len(board.Tiles[line]))
				for row := 0; row < 3; row++ {
					boardTile := board.Tiles[line][row]
					assert.NotNil(t, boardTile)
					assert.NotNil(t, boardTile.Tile)
				}
			}

			assert.NotNil(t, board.RemainingTile)
			assert.NotNil(t, board.RemainingTile.Tile)
		})
	})

	t.Run("Save()", func(t *testing.T) {
		var (
			testSaveId   = "test-save"
			testSavePath = store.filePath(testSaveId)
		)
		t.Run("Should save board in the given state file", func(t *testing.T) {
			t.Cleanup(func() {
				os.Remove(testSavePath)
			})

			// Save
			board, err := store.Load("test-board")
			assert.Nil(t, err)
			assert.Nil(t, store.Save(testSaveId, board))

			// Load and test
			stateJson, err := os.ReadFile(testSavePath)
			assert.Nil(t, err)

			savedBoard := new(model.Board)
			json.Unmarshal(stateJson, savedBoard)

			assert.Equal(t, board, savedBoard)
		})
	})
}
