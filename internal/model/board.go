package model

import (
	"errors"
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

	// Tiles are the tiles that are placed on a board.
	Tiles [][]*BoardTile

	// RemainingTile is the tile that was not placed on the board.
	RemainingTile *BoardTile
}

func (b *Board) InsertTileTopAt(row int) error {
	if (row & 1) != 1 {
		return errors.New("row must be odd")
	}

	var current = b.RemainingTile
	for line := 0; line < b.Size(); line++ {
		b.Tiles[line][row], current = current, b.Tiles[line][row]
	}

	b.RemainingTile = current
	return nil
}

func (b *Board) InsertTileRightAt(line int) error {
	if (line & 1) != 1 {
		return errors.New("row must be odd")
	}

	var current = b.RemainingTile
	for row := b.Size() - 1; row >= 0; row-- {
		b.Tiles[line][row], current = current, b.Tiles[line][row]
	}

	b.RemainingTile = current
	return nil
}

func (b *Board) InsertTileBottomAt(row int) error {
	if (row & 1) != 1 {
		return errors.New("row must be odd")
	}

	var current = b.RemainingTile
	for line := b.Size() - 1; line >= 0; line-- {
		b.Tiles[line][row], current = current, b.Tiles[line][row]
	}

	b.RemainingTile = current
	return nil
}

func (b *Board) InsertTileLeftAt(line int) error {
	if (line & 1) != 1 {
		return errors.New("row must be odd")
	}

	var current = b.RemainingTile
	for row := 0; row < b.Size(); row++ {
		b.Tiles[line][row], current = current, b.Tiles[line][row]
	}

	b.RemainingTile = current
	return nil
}

func (b *Board) RotateRemainingTileClockwise() {
	switch b.RemainingTile.Rotation {
	case Rotation0:
		b.RemainingTile.Rotation = Rotation90
	case Rotation90:
		b.RemainingTile.Rotation = Rotation180
	case Rotation180:
		b.RemainingTile.Rotation = Rotation270
	case Rotation270:
		b.RemainingTile.Rotation = Rotation0
	}
}

func (b *Board) RotateRemainingTileAntiClockwise() {
	switch b.RemainingTile.Rotation {
	case Rotation0:
		b.RemainingTile.Rotation = Rotation270
	case Rotation90:
		b.RemainingTile.Rotation = Rotation0
	case Rotation180:
		b.RemainingTile.Rotation = Rotation90
	case Rotation270:
		b.RemainingTile.Rotation = Rotation180
	}
}

// Size returns the board size in tiles.
func (b Board) Size() int {
	return len(b.Tiles)
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

	board.RemainingTile = &BoardTile{
		Tile:     tiles[len(tiles)-1],
		Rotation: Rotation0,
	}

	return board, nil
}