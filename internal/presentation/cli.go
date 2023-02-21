package presentation

import (
	"fmt"
	"io"

	"github.com/marmelab/labyrinth/internal/model"
)

const (
	// TileHeigth is the tile height in lines.
	TileHeigth = 3

	// TreasureRow is the row where the trasure is placed on a shape.
	// This is used to format a tile when printing it.
	TreasureRow = 1
)

// This is the shape parts for line assembly.
var tileShapeRows = map[model.Shape](map[model.Rotation](map[int]string)){
	model.ShapeI: map[model.Rotation](map[int]string){
		model.Rotation0: map[int]string{
			0: "───",
			1: ".%s.",
			2: "───",
		},
		model.Rotation90: map[int]string{
			0: "│.│",
			1: "│%s│",
			2: "│.│",
		},
		model.Rotation180: map[int]string{
			0: "───",
			1: ".%s.",
			2: "───",
		},
		model.Rotation270: map[int]string{
			0: "│.│",
			1: "│%s│",
			2: "│.│",
		},
	},
	model.ShapeT: map[model.Rotation](map[int]string){
		model.Rotation0: map[int]string{
			0: "┘.└",
			1: ".%s.",
			2: "───",
		},
		model.Rotation90: map[int]string{
			0: "│.└",
			1: "│%s.",
			2: "│.┌",
		},
		model.Rotation180: map[int]string{
			0: "───",
			1: ".%s.",
			2: "┐.┌",
		},
		model.Rotation270: map[int]string{
			0: "┘.│",
			1: ".%s│",
			2: "┐.│",
		},
	},
	model.ShapeV: map[model.Rotation](map[int]string){
		model.Rotation0: map[int]string{
			0: "──┐",
			1: ".%s│",
			2: "┐.│",
		},
		model.Rotation90: map[int]string{
			0: "┘.│",
			1: ".%s│",
			2: "──┘",
		},
		model.Rotation180: map[int]string{
			0: "│.└",
			1: "│%s.",
			2: "└──",
		},
		model.Rotation270: map[int]string{
			0: "┌──",
			1: "│%s.",
			2: "│.┌",
		},
	},
}

// BoardDrawer is in charge of drawing the board to the CLI.
type BoardDrawer interface {

	// DrawTo draws the board to the writer.
	DrawTo(w io.Writer, board *model.Board) (err error)
}

type boardDrawer struct {
}

// drawTileRow draws a tile row without a treasure.
func (d boardDrawer) drawTileRow(w io.Writer, boardTile *model.BoardTile, tileRow int) error {
	_, err := io.WriteString(w, tileShapeRows[boardTile.Tile.Shape][boardTile.Rotation][tileRow])
	return err
}

// drawTileRowWithTreasure draws a tile rows with a treasure in it.
func (d boardDrawer) drawTileRowWithTreasure(w io.Writer, boardTile *model.BoardTile) error {
	_, err := io.WriteString(w,
		fmt.Sprintf(
			tileShapeRows[boardTile.Tile.Shape][boardTile.Rotation][TreasureRow],
			string(boardTile.Tile.Treasure)))
	return err
}

func (d boardDrawer) DrawTo(w io.Writer, board *model.Board) (err error) {
	for _, line := range board.Tiles {
		for tileRow := 0; tileRow < TileHeigth; tileRow++ {
			for _, boardTile := range line {
				var err error = nil
				if tileRow != TreasureRow {
					err = d.drawTileRow(w, boardTile, tileRow)
				} else {
					err = d.drawTileRowWithTreasure(w, boardTile)
				}

				if err != nil {
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
