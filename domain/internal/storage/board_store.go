package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

var (
	// ErrMarshalFailure is returned when state file fails to be marshalled or
	// unmarshalled.
	ErrMarshalFailure = errors.New("failed to marshal/unmarshal board")
)

// BoardSaveFn is a function to save a board without relying on the save ID.
type BoardSaveFn func() (err error)

// BoardStore is in charge of storing board state.
type BoardStore interface {

	// Load the board with the given ID.
	Load(id string) (board *model.Board, err error)

	// Save the board.
	Save(id string, board *model.Board) (err error)

	// Get board using its ID, or create a new one if save does not exist yet.
	// This function also returns a save function that provide an  easy way to
	// store the save without knowing the save ID.
	//
	// Note:
	//   The board size and player count will be used to initialize the board if
	//   the board is not stored yet.
	Get(id string, boardSize, playerCount int) (board *model.Board, saveBoard BoardSaveFn, err error)
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

func (s fileBoardStore) Get(id string, boardSize, playerCount int) (*model.Board, BoardSaveFn, error) {
	board, err := s.Load(id)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Failed to load board: %v", err)
			return nil, nil, err
		}

		board, err = model.NewBoard(boardSize, playerCount)
		if err != nil {
			log.Printf("Failed to create board: %v", err)
			return nil, nil, err
		}
	}

	if err := s.Save(id, board); err != nil {
		log.Printf("Failed to save board: %v", err)
		return nil, nil, err
	}

	return board, func() error {
		if board.State == model.GameStateEnd {
			os.Remove(s.filePath(id))
			return nil
		}
		return s.Save(id, board)
	}, nil
}

// NewFileBoardStore returns a file-backed board store with the given root
// directory path.
func NewFileBoardStore(rootDirectory string) (BoardStore, error) {
	if err := os.MkdirAll(rootDirectory, 0755); err != nil {
		return nil, err
	}

	return fileBoardStore{
		rootDirectory: rootDirectory,
	}, nil
}
