package presentation

import (
	"fmt"
	"io"

	"github.com/marmelab/labyrinth/internal/model"
)

const (
	// TileHeigth is the tile height in lines.
	TileHeigth int = 3

	// UpperRow is the upper row of a tile
	UpperRow int = 0

	// TreasureRow is the row where the trasure is placed on a shape.
	// This is used to format a tile when printing it.
	TreasureRow int = 1

	// LowerRow is the lower row of a tile
	LowerRow int = 2
)

// BoardDrawer is in charge of drawing the board to the CLI.
type BoardDrawer interface {

	// DrawTo draws the board to the writer.
	DrawTo(w io.Writer, board *model.Board) (err error)
}

type boardDrawer struct {
}

func (d boardDrawer) write(w io.Writer, str string) error {
	_, err := io.WriteString(w, str)
	return err
}

func (d boardDrawer) drawITileRow(w io.Writer, rotation model.Rotation, tileRow int, treasure rune) error {

	if rotation == model.Rotation0 || rotation == model.Rotation180 {
		switch tileRow {
		case UpperRow, LowerRow:
			return d.write(w, "───")
		case TreasureRow:
			return d.write(w, "."+string(treasure)+".")
		default:
			return fmt.Errorf("unsupported row: %v", tileRow)
		}
	}

	if rotation == model.Rotation90 || rotation == model.Rotation270 {
		switch tileRow {
		case UpperRow, LowerRow:
			return d.write(w, "│.│")
		case TreasureRow:
			return d.write(w, "│"+string(treasure)+"│")
		default:
			return fmt.Errorf("unsupported row: %v", tileRow)
		}
	}

	return fmt.Errorf("unsupported rotation: %v", rotation)
}

func (d boardDrawer) drawTTileRow(w io.Writer, rotation model.Rotation, tileRow int, treasure rune) error {

	if rotation == model.Rotation0 {
		switch tileRow {
		case UpperRow:
			return d.write(w, "┘.└")
		case TreasureRow:
			return d.write(w, "."+string(treasure)+".")
		case LowerRow:
			return d.write(w, "───")
		default:
			return fmt.Errorf("unsupported row: %v", tileRow)
		}
	}

	if rotation == model.Rotation90 {
		switch tileRow {
		case UpperRow:
			return d.write(w, "│.└")
		case TreasureRow:
			return d.write(w, "│"+string(treasure)+".")
		case LowerRow:
			return d.write(w, "│.┌")
		default:
			return fmt.Errorf("unsupported row: %v", tileRow)
		}
	}

	if rotation == model.Rotation180 {
		switch tileRow {
		case UpperRow:
			return d.write(w, "───")
		case TreasureRow:
			return d.write(w, "."+string(treasure)+".")
		case LowerRow:
			return d.write(w, "┐.┌")
		default:
			return fmt.Errorf("unsupported row: %v", tileRow)
		}
	}

	if rotation == model.Rotation270 {
		switch tileRow {
		case UpperRow:
			return d.write(w, "┘.│")
		case TreasureRow:
			return d.write(w, "."+string(treasure)+"│")
		case LowerRow:
			return d.write(w, "┐.│")
		default:
			return fmt.Errorf("unsupported row: %v", tileRow)
		}
	}

	return fmt.Errorf("unsupported rotation: %v", rotation)
}

func (d boardDrawer) drawVTileRow(w io.Writer, rotation model.Rotation, tileRow int, treasure rune) error {

	if rotation == model.Rotation0 {
		switch tileRow {
		case UpperRow:
			return d.write(w, "──┐")
		case TreasureRow:
			return d.write(w, "."+string(treasure)+"│")
		case LowerRow:
			return d.write(w, "┐.│")
		default:
			return fmt.Errorf("unsupported row: %v", tileRow)
		}
	}

	if rotation == model.Rotation90 {
		switch tileRow {
		case UpperRow:
			return d.write(w, "┘.│")
		case TreasureRow:
			return d.write(w, "."+string(treasure)+"│")
		case LowerRow:
			return d.write(w, "──┘")
		default:
			return fmt.Errorf("unsupported row: %v", tileRow)
		}
	}

	if rotation == model.Rotation180 {
		switch tileRow {
		case UpperRow:
			return d.write(w, "│.└")
		case TreasureRow:
			return d.write(w, "│"+string(treasure)+".")
		case LowerRow:
			return d.write(w, "└──")
		default:
			return fmt.Errorf("unsupported row: %v", tileRow)
		}
	}

	if rotation == model.Rotation270 {
		switch tileRow {
		case UpperRow:
			return d.write(w, "┌──")
		case TreasureRow:
			return d.write(w, "│"+string(treasure)+".")
		case LowerRow:
			return d.write(w, "│.┌")
		default:
			return fmt.Errorf("unsupported row: %v", tileRow)
		}
	}

	return fmt.Errorf("unsupported rotation: %v", rotation)
}

// drawTileRow draws a tile row without a treasure.
func (d boardDrawer) drawTileRow(w io.Writer, boardTile *model.BoardTile, tileRow int) error {
	switch boardTile.Tile.Shape {
	case model.ShapeI:
		return d.drawITileRow(w, boardTile.Rotation, tileRow, boardTile.Tile.Treasure)

	case model.ShapeT:
		return d.drawTTileRow(w, boardTile.Rotation, tileRow, boardTile.Tile.Treasure)

	case model.ShapeV:
		return d.drawVTileRow(w, boardTile.Rotation, tileRow, boardTile.Tile.Treasure)
	}

	return fmt.Errorf("unsupported tile shape : %v", boardTile.Tile.Shape)
}

func (d boardDrawer) DrawTo(w io.Writer, board *model.Board) (err error) {
	for _, line := range board.Tiles {
		for tileRow := 0; tileRow < TileHeigth; tileRow++ {
			for _, boardTile := range line {
				if err := d.drawTileRow(w, boardTile, tileRow); err != nil {
					return err
				}
			}

			if _, err := w.Write([]byte{'\n'}); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewBoardDrawer() BoardDrawer {
	return &boardDrawer{}
}
