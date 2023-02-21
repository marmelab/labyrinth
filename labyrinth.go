package main

import (
	"log"

	"github.com/marmelab/labyrinth/internal/model"
	"github.com/marmelab/labyrinth/internal/presentation"
)

func main() {
	board, err := model.NewBoard(7)
	if err != nil {
		log.Fatalf("Failed to initialize board: %v.", err)
	}

	if err := presentation.GameLoop(board); err != nil {
		log.Fatalf("Failed to run main loop: %v.", err)
	}
}
