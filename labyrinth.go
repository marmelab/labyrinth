package main

import (
	"log"
	"os"
	"path"

	"github.com/marmelab/labyrinth/internal/presentation"
	"github.com/marmelab/labyrinth/internal/storage"
	"github.com/spf13/cobra"
)

var (
	// LoadStateId holds the state to load from file
	LoadStateId string

	// BoardSize holds the board size param.
	BoardSize int

	// PlayerCount holds the player count param.
	PlayerCount int
)

var (
	homeDirectory, _ = os.UserHomeDir()
	storageDirectory = path.Join(homeDirectory, ".labyrinth")
)

func RunApplication(cmd *cobra.Command, args []string) {
	boardStore, err := storage.NewFileBoardStore(storageDirectory)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v.", err)
	}

	board, saveBoard, err := boardStore.Get(LoadStateId, BoardSize, PlayerCount)
	if err != nil {
		log.Fatalf("Failed to initialize board: %v.", err)
	}

	if err := presentation.RunGameLoop(board, saveBoard); err != nil {
		log.Fatalf("Failed to run main loop: %v.", err)
	}
}

func main() {
	var command = &cobra.Command{
		Use:   "labyrinth",
		Short: "Play the labyrinth game on you CLI",
		Run:   RunApplication,
	}

	command.
		PersistentFlags().
		StringVarP(&LoadStateId, "save", "s", "_default", "The save to load, it will be created if it does not exist yet.")

	command.
		PersistentFlags().
		IntVarP(&BoardSize, "size", "b", 7, "The board size when initialized, defaults to 7.")

	command.
		PersistentFlags().
		IntVarP(&PlayerCount, "players", "p", 1, "The number of players, defaults to 1.")

	if err := command.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v.", err)
	}
}
