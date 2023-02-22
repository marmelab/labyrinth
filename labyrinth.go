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

	board, saveBoard, err := boardStore.Get(LoadStateId)
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
		StringVarP(&LoadStateId, "save", "s", "_default", "the save to load, it will be created if it does not exist yet.")

	if err := command.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v.", err)
	}
}
