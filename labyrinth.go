package main

import (
	"log"
	"os"

	"github.com/marmelab/labyrinth/internal/model"
	"github.com/marmelab/labyrinth/internal/presentation"
)

func main() {
	board, err := model.NewBoard(3)
	if err != nil {
		log.Fatalf("Failed to initialize board: %v.", err)
	}

	drawer := presentation.NewBoardDrawer()

	drawer.DrawTo(os.Stdout, board)
}
