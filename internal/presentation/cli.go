package presentation

import (
	"fmt"
	"io"

	"github.com/marmelab/labyrinth/internal/model"
)

// This is the shape parts for line assembly
var shapesParts = map[model.Shape](map[model.Rotation](map[int]string)){
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

type BoardDrawer interface {

	// DrawTo draws the board to the writer.
	DrawTo(w io.Writer, board *model.Board) (err error)
}

type boardDrawer struct {
}

func (d boardDrawer) DrawTo(w io.Writer, board *model.Board) (err error) {
	for _, line := range board.Tiles {
		for i := 0; i < 3; i++ {
			for _, bt := range line {
				var bytes []byte
				if i != 1 {
					bytes = []byte(shapesParts[bt.Tile.Shape][bt.Rotation][i])
				} else {
					bytes = []byte(fmt.Sprintf(shapesParts[bt.Tile.Shape][bt.Rotation][i], string(bt.Tile.Treasure)))
				}

				if _, err := w.Write(bytes); err != nil {
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
