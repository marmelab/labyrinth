package presentation

import (
	"io"

	"github.com/marmelab/labyrinth/internal/model"
)

const (
	// TileSize is the tile height in lines.
	TileSize int = 3

	// TileBorderSize is the tile border size in lines.
	TileBorderSize = 1

	// TileOuterSize is the tile height in lines including border.
	TileOuterSize int = TileSize + TileBorderSize

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

	// DrawBoardTileTo draws a board tile to the writer.
	DrawBoardTileTo(w io.Writer, boardTile *model.BoardTile) (err error)
}

type boardDrawer struct {
}

func (d boardDrawer) writeTile(buffer, tile [][]rune, line, row int) {
	for i := 0; i < TileSize; i++ {
		var (
			bufferX = line*TileSize + i
			bufferY = row * TileSize
		)
		copy(buffer[bufferX][bufferY:], tile[i])
	}
}

func (d boardDrawer) formatITile(rotation model.Rotation, treasure model.Treasure) [][]rune {
	tre := rune(treasure)

	switch rotation {
	case model.Rotation0, model.Rotation180:
		return [][]rune{
			{'─', '─', '─'},
			{' ', tre, ' '},
			{'─', '─', '─'},
		}
	default:
		return [][]rune{
			{'│', ' ', '│'},
			{'│', tre, '│'},
			{'│', ' ', '│'},
		}
	}
}

func (d boardDrawer) formatTTile(rotation model.Rotation, treasure model.Treasure) [][]rune {
	tre := rune(treasure)

	switch rotation {
	case model.Rotation0:
		return [][]rune{
			{'┘', ' ', '└'},
			{' ', tre, ' '},
			{'─', '─', '─'},
		}
	case model.Rotation90:
		return [][]rune{
			{'│', ' ', '└'},
			{'│', tre, ' '},
			{'│', ' ', '┌'},
		}
	case model.Rotation180:
		return [][]rune{
			{'─', '─', '─'},
			{' ', tre, ' '},
			{'┐', ' ', '┌'},
		}
	default:
		return [][]rune{
			{'┘', ' ', '│'},
			{' ', tre, '│'},
			{'┐', ' ', '│'},
		}
	}
}

func (d boardDrawer) formatVTile(rotation model.Rotation, treasure model.Treasure) [][]rune {
	tre := rune(treasure)

	switch rotation {
	case model.Rotation0:
		return [][]rune{
			{'─', '─', '┐'},
			{' ', tre, '│'},
			{'┐', ' ', '│'},
		}
	case model.Rotation90:
		return [][]rune{
			{'┘', ' ', '│'},
			{' ', tre, '│'},
			{'─', '─', '┘'},
		}
	case model.Rotation180:
		return [][]rune{
			{'│', ' ', '└'},
			{'│', tre, ' '},
			{'└', '─', '─'},
		}
	default:
		return [][]rune{
			{'┌', '─', '─'},
			{'│', tre, ' '},
			{'│', ' ', '┌'},
		}
	}
}

func (d boardDrawer) drawTile(buffer [][]rune, boardTile *model.BoardTile, line, row int) {
	switch boardTile.Tile.Shape {
	case model.ShapeI:
		d.writeTile(buffer, d.formatITile(boardTile.Rotation, boardTile.Tile.Treasure), line, row)

	case model.ShapeT:
		d.writeTile(buffer, d.formatTTile(boardTile.Rotation, boardTile.Tile.Treasure), line, row)

	case model.ShapeV:
		d.writeTile(buffer, d.formatVTile(boardTile.Rotation, boardTile.Tile.Treasure), line, row)
	}
}

func (d boardDrawer) initBuffer(tileCount int) [][]rune {
	var (
		bufferSize = TileSize * tileCount
		buffer     = make([][]rune, bufferSize)
	)
	for i := 0; i < bufferSize; i++ {
		buffer[i] = make([]rune, bufferSize)
	}
	return buffer
}

func (d boardDrawer) DrawTo(w io.Writer, board *model.Board) (err error) {
	buffer := d.initBuffer(len(board.Tiles))

	for i, line := range board.Tiles {
		for j, boardTile := range line {
			d.drawTile(buffer, boardTile, i, j)
		}

		for outputLine := 0; outputLine < TileSize; outputLine++ {
			io.WriteString(w, string(buffer[i*TileSize+outputLine])+"\n")
		}
	}
	return nil
}

func (d boardDrawer) DrawBoardTileTo(w io.Writer, boardTile *model.BoardTile) (err error) {
	buffer := d.initBuffer(1)

	d.drawTile(buffer, boardTile, 0, 0)
	for outputLine := 0; outputLine < TileSize; outputLine++ {
		io.WriteString(w, string(buffer[outputLine])+"\n")
	}

	return nil
}

func NewBoardDrawer() BoardDrawer {
	return &boardDrawer{}
}
