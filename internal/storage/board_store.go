package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"

	"github.com/marmelab/labyrinth/internal/model"
)

var (
	// ErrMarshalFailure is returned when state file fails to be marshalled or
	// unmarshalled.
	ErrMarshalFailure = errors.New("failed to marshal/unmarshal board")
)

// BoardStore is in charge of storing board state.
type BoardStore interface {

	// Load the board with the given ID.
	Load(id string) (board *model.Board, err error)

	// Save the board.
	Save(id string, board *model.Board) (err error)
}

// fileBoardStore is a board store that uses a file store
type fileBoardStore struct {
	rootDirectory string
}

func (s fileBoardStore) filePath(id string) string {
	return path.Join(s.rootDirectory, id)
}

func (s fileBoardStore) Load(id string) (*model.Board, error) {
	stateJson, err := os.ReadFile(s.filePath(id))
	if err != nil {
		return nil, err
	}

	board := new(model.Board)
	if err := json.Unmarshal(stateJson, board); err != nil {
		log.Printf("Failed to unmarshal state: %v.", err)
		return nil, ErrMarshalFailure
	}

	return board, nil
}

func (s fileBoardStore) Save(id string, board *model.Board) error {
	stateJson, err := json.Marshal(board)
	if err != nil {
		log.Printf("Failed to marshal state: %v.", err)
		return ErrMarshalFailure
	}

	if err := os.WriteFile(s.filePath(id), stateJson, 0644); err != nil {
		return err
	}

	return nil
}

// NewFileBoardStore returns a file-backed board store with the given root
// directory path.
func NewFileBoardStore(rootDirectory string) BoardStore {
	return fileBoardStore{
		rootDirectory: rootDirectory,
	}
}
