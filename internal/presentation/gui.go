package presentation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/marmelab/labyrinth/internal/model"
)

type gameUi struct {
	gui  *gocui.Gui
	loop *gameLoop

	boardDrawer BoardDrawer
	state       *model.Board
}

func (g gameUi) drawButton(name, text string, topLeftX, topLeftY, bottomRightX, bottomRightY int, handler GuiHandler) error {
	button, err := g.gui.SetView(name, topLeftX, topLeftY, bottomRightX, bottomRightY, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return errors.Join(fmt.Errorf("failed to initialize button %s", name), err)
		}

		button.BgColor = gocui.ColorGreen
		button.FgColor = gocui.ColorBlack
		button.Frame = false

		fmt.Fprint(button, text)
		if err := g.gui.SetKeybinding(name, gocui.MouseLeft, gocui.ModNone, handler); err != nil {
			return errors.Join(fmt.Errorf("failed to set mouse click to %s", name), err)
		}
	}

	return nil
}

func (g gameUi) drawBoard(tileCount, boardSize int) error {
	boardView, err := g.gui.SetView("board", BoardMargin, BoardMargin, boardSize+BoardMargin, boardSize+BoardMargin, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return errors.Join(errors.New("failed to initialize board"), err)
		}

		boardView.FgColor = gocui.ColorCyan
		for line := 0; line < boardSize; line++ {
			fmt.Fprintln(boardView, strings.Repeat("·", boardSize))
		}
	}

	if err := g.drawTiles(tileCount); err != nil {
		return errors.Join(errors.New("failed to initialize tiles"), err)
	}

	return nil
}

func (g gameUi) drawTiles(tileCount int) error {
	for line := 0; line < tileCount; line++ {
		for row := 0; row < tileCount; row++ {
			var (
				name         = fmt.Sprintf("tile-%d-%d", line, row)
				topLeftX     = BoardMargin + row*TileOuterSize + TileBorderSize
				topLeftY     = BoardMargin + line*TileOuterSize + TileBorderSize
				bottomRightX = topLeftX + TileSize + TileBorderSize
				bottomRightY = topLeftY + TileSize + TileBorderSize
			)

			tileView, err := g.gui.SetView(name, topLeftX, topLeftY, bottomRightX, bottomRightY, 1)
			if err != nil {
				if err != gocui.ErrUnknownView {
					return errors.Join(errors.New("failed to initialize board"), err)
				}
				tileView.Frame = false
			}

			tileView.Clear()
			if err := g.boardDrawer.DrawBoardTileTo(tileView, g.state.Tiles[line][row]); err != nil {
				return errors.Join(errors.New("failed to draw tile"), err)
			}
		}
	}
	return nil
}

func (g gameUi) drawBoardActions(tileCount, boardSize int) error {
	for buttonIndex := 1; buttonIndex < tileCount; buttonIndex += 2 {
		{
			var (
				name         = fmt.Sprintf("button-top-%d", buttonIndex)
				topLeftX     = BoardMargin + (TileOuterSize * buttonIndex) + 1
				topLeftY     = 0
				bottomRightX = topLeftX + TileOuterSize
				bottomRightY = topLeftY + TileOuterSize
			)

			if err := g.drawButton(name, "\n\n ↓", topLeftX, topLeftY, bottomRightX, bottomRightY, g.loop.insertTile(TopButton, buttonIndex)); err != nil {
				return err
			}
		}

		{
			var (
				name         = fmt.Sprintf("button-right-%d", buttonIndex)
				topLeftX     = BoardMargin + (boardSize + 1)
				topLeftY     = BoardMargin + (TileOuterSize * buttonIndex) + 1
				bottomRightX = topLeftX + TileOuterSize
				bottomRightY = topLeftY + TileOuterSize
			)

			if err := g.drawButton(name, "\n←", topLeftX, topLeftY, bottomRightX, bottomRightY, g.loop.insertTile(RightButton, buttonIndex)); err != nil {
				return err
			}
		}

		{
			var (
				name         = fmt.Sprintf("button-bottom-%d", buttonIndex)
				topLeftX     = BoardMargin + (TileOuterSize * buttonIndex) + 1
				topLeftY     = BoardMargin + (boardSize + 1)
				bottomRightX = topLeftX + TileOuterSize
				bottomRightY = topLeftY + TileOuterSize
			)

			if err := g.drawButton(name, " ↑", topLeftX, topLeftY, bottomRightX, bottomRightY, g.loop.insertTile(BottomButton, buttonIndex)); err != nil {
				return err
			}
		}

		{
			var (
				name         = fmt.Sprintf("button-left-%d", buttonIndex)
				topLeftX     = 0
				topLeftY     = BoardMargin + (TileOuterSize * buttonIndex) + 1
				bottomRightX = topLeftX + TileOuterSize
				bottomRightY = topLeftY + TileOuterSize
			)

			if err := g.drawButton(name, "\n  →", topLeftX, topLeftY, bottomRightX, bottomRightY, g.loop.insertTile(LeftButton, buttonIndex)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g gameUi) drawRemainingTile(boardSize, boardOffset int) error {

	var (
		topLeftX     = BoardMargin
		topLeftY     = BoardMargin + boardOffset
		bottomRightX = topLeftX + boardSize
		bottomRightY = topLeftY + 8
	)

	if _, err := g.gui.SetView("remaining-tile-box", topLeftX, topLeftY, bottomRightX, bottomRightY, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.Join(errors.New("failed to initialize remaining tile box"), err)
		}
	}

	{
		var (
			name         = "button-rotate-anticlockwise"
			topLeftX     = topLeftX + 5
			topLeftY     = topLeftY + 2
			bottomRightX = topLeftX + TileOuterSize
			bottomRightY = topLeftY + TileOuterSize
		)

		if err := g.drawButton(name, "\n ↶", topLeftX, topLeftY, bottomRightX, bottomRightY, g.loop.rotateRemainingTile(RotateAntiClockWise)); err != nil {
			return err
		}
	}

	{
		var (
			topLeftX     = topLeftX + 13
			topLeftY     = topLeftY + 2
			bottomRightX = topLeftX + 4
			bottomRightY = topLeftY + TileOuterSize
		)

		remainingTileView, err := g.gui.SetView(
			"remaining-tile",
			topLeftX,
			topLeftY,
			bottomRightX,
			bottomRightY, 1)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return errors.Join(errors.New("failed to initialize remaining tile view"), err)
			}
			remainingTileView.Frame = false
		}

		remainingTileView.Clear()
		if err := g.boardDrawer.DrawBoardTileTo(remainingTileView, g.state.RemainingTile); err != nil {
			return errors.Join(errors.New("failed to draw tile"), err)
		}
	}

	{
		var (
			name         = "button-rotate-clockwise"
			topLeftX     = topLeftX + 21
			topLeftY     = topLeftY + 2
			bottomRightX = topLeftX + 4
			bottomRightY = topLeftY + TileOuterSize
		)

		if err := g.drawButton(name, "\n ↷", topLeftX, topLeftY, bottomRightX, bottomRightY, g.loop.rotateRemainingTile(RotateClockWise)); err != nil {
			return err
		}
	}

	return nil
}

func (g gameUi) layout(gui *gocui.Gui) error {
	var (
		tileCount   = g.state.Size()
		boardSize   = tileCount*TileOuterSize + TileBorderSize + 1
		boardOffset = boardSize + TileOuterSize + 1
	)

	if err := g.drawBoard(tileCount, boardSize); err != nil {
		return err
	}

	if err := g.drawBoardActions(tileCount, boardSize); err != nil {
		return err
	}

	if err := g.drawRemainingTile(boardSize, boardOffset); err != nil {
		return err
	}

	return nil
}
