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
	board       *model.Board
}

func (g gameUi) backgroundColor() gocui.Attribute {
	switch g.board.GetCurrentPlayer().Color {
	case model.ColorBlue:
		return gocui.ColorBlue
	case model.ColorGreen:
		return gocui.ColorGreen
	case model.ColorRed:
		return gocui.ColorRed
	default:
		return gocui.ColorYellow
	}
}

func (g gameUi) foregroundColor() gocui.Attribute {
	switch g.board.GetCurrentPlayer().Color {
	case model.ColorBlue, model.ColorGreen, model.ColorRed:
		return gocui.ColorWhite
	default:
		return gocui.ColorBlack
	}
}

func (g gameUi) drawButton(name, text string, topLeftX, topLeftY, bottomRightX, bottomRightY int, handler GuiHandler) (*gocui.View, error) {
	button, err := g.gui.SetView(name, topLeftX, topLeftY, bottomRightX, bottomRightY, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return nil, errors.Join(fmt.Errorf("failed to initialize button %s", name), err)
		}

		button.Frame = false

		fmt.Fprint(button, text)
		if err := g.gui.SetKeybinding(name, gocui.MouseLeft, gocui.ModNone, handler); err != nil {
			return nil, errors.Join(fmt.Errorf("failed to set mouse click to %s", name), err)
		}
	}

	return button, nil
}

func (g gameUi) drawInsertTileButton(name, text string, topLeftX, topLeftY, bottomRightX, bottomRightY int, buttonType Button, buttonIndex int) error {
	button, err := g.drawButton(name, text, topLeftX, topLeftY, bottomRightX, bottomRightY, g.loop.insertTile(buttonType, buttonIndex))
	if err != nil {
		return err
	}

	if g.board.State == model.GameStatePlaceTile {
		button.BgColor = gocui.ColorMagenta
		button.FgColor = gocui.ColorWhite
	} else {
		button.BgColor = gocui.ColorDefault
		button.FgColor = gocui.ColorDefault
	}

	return nil
}

func (g gameUi) drawRotateTileButton(name, text string, topLeftX, topLeftY, bottomRightX, bottomRightY int, rotationType RotationType) error {
	button, err := g.drawButton(name, text, topLeftX, topLeftY, bottomRightX, bottomRightY, g.loop.rotateRemainingTile(rotationType))
	if err != nil {
		return err
	}

	button.BgColor = gocui.ColorMagenta
	button.FgColor = gocui.ColorWhite

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
	accessibleTiles := g.board.GetAccessibleTiles()

	for line := 0; line < tileCount; line++ {
		for row := 0; row < tileCount; row++ {
			var (
				name         = fmt.Sprintf("tile-%d-%d", line, row)
				topLeftX     = BoardMargin + row*TileOuterSize + TileBorderSize
				topLeftY     = BoardMargin + line*TileOuterSize + TileBorderSize
				bottomRightX = topLeftX + TileSize + TileBorderSize
				bottomRightY = topLeftY + TileSize + TileBorderSize

				currentTile = g.board.Tiles[line][row]
			)

			tileView, err := g.gui.SetView(name, topLeftX, topLeftY, bottomRightX, bottomRightY, 1)
			if err != nil {
				if err != gocui.ErrUnknownView {
					return errors.Join(errors.New("failed to initialize board"), err)
				}
				tileView.Frame = false

				if err := g.gui.SetKeybinding(name, gocui.MouseLeft, gocui.ModNone, g.loop.moveCurrentPlayerTo(line, row)); err != nil {
					return errors.Join(fmt.Errorf("failed to set mouse click to %s", name), err)
				}
			}

			tileView.Clear()
			if err := g.boardDrawer.DrawBoardTileTo(tileView, g.board.Tiles[line][row]); err != nil {
				return errors.Join(errors.New("failed to draw tile"), err)
			}

			var (
				currentPlayer = g.board.GetCurrentPlayer()
				targets       = currentPlayer.Targets
			)

			if currentPlayer.Position.Line == line && currentPlayer.Position.Row == row {
				tileView.BgColor = g.backgroundColor()
				tileView.FgColor = g.foregroundColor()
			} else if len(targets) > 0 && targets[0] == currentTile.Tile.Treasure {
				tileView.BgColor = gocui.ColorCyan
				tileView.FgColor = gocui.ColorBlack
			} else if g.board.State == model.GameStateMovePawn &&
				accessibleTiles.Contains(&model.Coordinate{Line: line, Row: row}) {

				tileView.BgColor = gocui.ColorMagenta
				tileView.FgColor = gocui.ColorWhite
			} else {
				tileView.BgColor = gocui.ColorWhite
				tileView.FgColor = gocui.ColorBlack
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

			if err := g.drawInsertTileButton(name, "\n ↓", topLeftX, topLeftY, bottomRightX, bottomRightY, TopButton, buttonIndex); err != nil {
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

			if err := g.drawInsertTileButton(name, "\n ←", topLeftX, topLeftY, bottomRightX, bottomRightY, RightButton, buttonIndex); err != nil {
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

			if err := g.drawInsertTileButton(name, "\n ↑", topLeftX, topLeftY, bottomRightX, bottomRightY, BottomButton, buttonIndex); err != nil {
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

			if err := g.drawInsertTileButton(name, "\n →", topLeftX, topLeftY, bottomRightX, bottomRightY, LeftButton, buttonIndex); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g gameUi) drawRemainingTile(leftOffset int) error {
	var (
		boardSize   = 7*TileOuterSize + TileBorderSize + 1
		boardOffset = boardSize + TileOuterSize + 1
	)

	var (
		topLeftX     = leftOffset + BoardMargin + 1
		topLeftY     = boardOffset - 3
		bottomRightX = topLeftX + 27
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
			topLeftX     = topLeftX + 3
			topLeftY     = topLeftY + 2
			bottomRightX = topLeftX + TileOuterSize
			bottomRightY = topLeftY + TileOuterSize
		)

		if err := g.drawRotateTileButton(name, "\n ↶", topLeftX, topLeftY, bottomRightX, bottomRightY, RotateAntiClockWise); err != nil {
			return err
		}
	}

	{
		var (
			topLeftX     = topLeftX + 11
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
		if err := g.boardDrawer.DrawBoardTileTo(remainingTileView, g.board.RemainingTile); err != nil {
			return errors.Join(errors.New("failed to draw tile"), err)
		}
	}

	{
		var (
			name         = "button-rotate-clockwise"
			topLeftX     = topLeftX + 19
			topLeftY     = topLeftY + 2
			bottomRightX = topLeftX + 4
			bottomRightY = topLeftY + TileOuterSize
		)

		if err := g.drawRotateTileButton(name, "\n ↷", topLeftX, topLeftY, bottomRightX, bottomRightY, RotateClockWise); err != nil {
			return err
		}
	}

	return nil
}

func (g gameUi) drawCurrentPlayer(boardOffset int) error {

	var (
		topLeftX     = BoardMargin + boardOffset + 1
		topLeftY     = BoardMargin
		bottomRightX = topLeftX + 27
		bottomRightY = topLeftY + 6
	)

	currentPlayerBox, err := g.gui.SetView("current-player-box", topLeftX, topLeftY, bottomRightX, bottomRightY, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return errors.Join(errors.New("failed to initialize current player box"), err)
		}
	}

	currentPlayerBox.Clear()

	if g.board.State == model.GameStateEnd {
		fmt.Fprintf(currentPlayerBox, `

         GAME OVER
`)
		return nil
	}

	currentPlayer := g.board.GetCurrentPlayer()
	fmt.Fprintf(currentPlayerBox, `
Current player: %10s

Target: %18s
`, currentPlayer.Name(), string(currentPlayer.Targets[0]))

	return nil
}

func (g gameUi) drawScores(boardOffset int) error {

	var (
		topLeftX     = BoardMargin + boardOffset + 1
		topLeftY     = BoardMargin + 8
		bottomRightX = topLeftX + 27
		bottomRightY = topLeftY + 2*len(g.board.Players) + 4
	)
	scoreBox, err := g.gui.SetView("score-box", topLeftX, topLeftY, bottomRightX, bottomRightY, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return errors.Join(errors.New("failed to initialize current player box"), err)
		}
	}

	scoreBox.Clear()
	fmt.Fprintf(scoreBox, `
Scores:
`)
	for _, player := range g.board.Players {
		fmt.Fprintf(scoreBox, `
%-10s %15d
`, player.Name(), player.Score)
	}

	return nil
}

func (g gameUi) layout(gui *gocui.Gui) error {
	var (
		tileCount   = g.board.GetSize()
		boardSize   = tileCount*TileOuterSize + TileBorderSize + 1
		boardOffset = boardSize + TileOuterSize + 1
	)

	if err := g.drawBoard(tileCount, boardSize); err != nil {
		return err
	}

	if err := g.drawBoardActions(tileCount, boardSize); err != nil {
		return err
	}

	if err := g.drawRemainingTile(boardOffset); err != nil {
		return err
	}

	if err := g.drawCurrentPlayer(boardOffset); err != nil {
		return err
	}

	if err := g.drawScores(boardOffset); err != nil {
		return err
	}
	return nil
}
