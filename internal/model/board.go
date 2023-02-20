package model

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	IShapedPercentage                = 0.26
	TShapedPercentage                = 0.36
	VShapedWithoutTreasurePercentage = 0.20
	VShapedWithTreasurePercentage    = 0.18
)

var (
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Board represents the game board.
type Board struct {

	// Tiles are the tile sthat are placed on a board.
	Tiles [][]*BoardTile

	// RemainingTile is the tile that was not placed on the board.
	RemainingTile *Tile
}

// BoardTile represents a tile that is placed on a board with a given rotation.
type BoardTile struct {

	// Tile is the underlying tile.
	Tile *Tile

	// Rotation is the tile rotation.
	Rotation Rotation
}

// Rotation represents a tile rotation on a board.
type Rotation int

const (
	Rotation0 Rotation = 0

	Rotation90 Rotation = 90

	Rotation180 Rotation = 180

	Rotation270 Rotation = 270
)

func generateTiles(size int) (tiles []*Tile, treasureCount int) {
	var (
		tileCount                = size*size + 1
		tShapedThreshold         = int(math.Round(TShapedPercentage * float64(tileCount)))
		vShapedWithTreasureCount = tShapedThreshold + int(math.Round(VShapedWithTreasurePercentage*float64(tileCount)))
		iShapedThreshold         = vShapedWithTreasureCount + int(math.Round(IShapedPercentage*float64(tileCount)))
	)

	tiles = make([]*Tile, 0, tileCount)

	for i := 0; i < tileCount; i++ {
		if i < tShapedThreshold {
			tiles = append(tiles, &Tile{
				Shape:    ShapeT,
				Treasure: 'A' + rune(i),
			})
		} else if i < vShapedWithTreasureCount {
			tiles = append(tiles, &Tile{
				Shape:    ShapeV,
				Treasure: 'A' + rune(i),
			})
		} else if i < iShapedThreshold {
			tiles = append(tiles, &Tile{
				Shape:    ShapeI,
				Treasure: NoTreasure,
			})
		} else {
			tiles = append(tiles, &Tile{
				Shape:    ShapeV,
				Treasure: NoTreasure,
			})
		}
	}

	return tiles, vShapedWithTreasureCount
}

func randomRotation() Rotation {
	switch randomGenerator.Int63n(4) {
	case 0:
		return Rotation0
	case 1:
		return Rotation90
	case 2:
		return Rotation90
	default:
		return Rotation270
	}
}

// NewBoard returns a board for the given size.
// The size param MUST be an odd number.
func NewBoard(size int) (*Board, error) {
	if (size & 1) != 1 {
		return nil, fmt.Errorf("size must be an odd number, got: %d", size)
	}

	tiles, _ := generateTiles(size)

	randomGenerator.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	board := &Board{
		Tiles: make([][]*BoardTile, size),
	}

	for i := 0; i < size; i++ {
		board.Tiles[i] = make([]*BoardTile, size)
		for j := 0; j < size; j++ {
			board.Tiles[i][j] = &BoardTile{
				Tile:     tiles[(size*i)+j],
				Rotation: randomRotation(),
			}
		}
	}

	board.RemainingTile = tiles[len(tiles)-1]

	return board, nil
}
