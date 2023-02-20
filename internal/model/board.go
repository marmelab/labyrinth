package model

import (
	"fmt"
)

const (
	IShapedPercentage               = 0.26
	TShapedPercentage               = 0.36
	VShapedWithouTreasurePercentage = 0.20
	VShapedWithTreasurePercentage   = 0.18
)

// Board represents the game board.
type Board struct {

	// Tiles are the tile sthat are placed on a board.
	Tiles [][]BoardTile
}

// BoardTile represents a tile that is placed on a board with a given rotation.
type BoardTile struct {

	// Tile is the underlying tile
	Tile Tile

	// Rotation is the tile rotation
	Rotation Rotation
}

// Rotation represents a tile rotation on a board.
type Rotation int

const (
	Rotation000 Rotation = 0

	Rotation090 Rotation = 90

	Rotation180 Rotation = 180

	Rotation270 Rotation = 270
)

// NewBoard returns a board for the given size.
// The size param MUST be an odd number.
func NewBoard(size int) (*Board, error) {
	if (size & 1) != 1 {
		return nil, fmt.Errorf("size must be an odd number, got: %d", size)
	}

	return &Board{}, nil
}
