package presentation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
	"github.com/marmelab/labyrinth/internal/model"
)

const (
	BoardMargin = 4
)

type Button int

const (
	TopButton Button = iota
	RightButton
	BottomButton
	LeftButton
)

type RotationType int

const (
	RotateClockWise RotationType = iota
	RotateAntiClockWise
)

type GuiHandler func(gui *gocui.Gui, view *gocui.View) error

type gameLoop struct {
	boardDrawer BoardDrawer
	state       *model.Board

	gui *gocui.Gui
}

func (g gameLoop) drawButton(name, text string, topLeftX, topLeftY, bottomRightX, bottomRightY int, handler GuiHandler) error {
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

func (g gameLoop) drawBoard(tileCount, boardSize int) error {
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

func (g gameLoop) drawTiles(tileCount int) error {
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

func (g gameLoop) drawBoardActions(tileCount, boardSize int) error {
	for buttonIndex := 1; buttonIndex < tileCount; buttonIndex += 2 {
		{
			var (
				name         = fmt.Sprintf("button-top-%d", buttonIndex)
				topLeftX     = BoardMargin + (TileOuterSize * buttonIndex) + 1
				topLeftY     = 0
				bottomRightX = topLeftX + TileOuterSize
				bottomRightY = topLeftY + TileOuterSize
			)

			if err := g.drawButton(name, "\n ↓", topLeftX, topLeftY, bottomRightX, bottomRightY, g.insertTile(TopButton, buttonIndex)); err != nil {
				return err
			}
		}

		{
			var (
				name         = fmt.Sprintf("button-right-%d", buttonIndex)
				topLeftX     = BoardMargin + (boardSize + 1)
				topLeftY     = BoardMargin + (TileOuterSize * buttonIndex) + 1
				bottomRightX = topLeftX + TileSize
				bottomRightY = topLeftY + TileOuterSize
			)

			if err := g.drawButton(name, "\n←", topLeftX, topLeftY, bottomRightX, bottomRightY, g.insertTile(RightButton, buttonIndex)); err != nil {
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

			if err := g.drawButton(name, " ↑", topLeftX, topLeftY, bottomRightX, bottomRightY, g.insertTile(BottomButton, buttonIndex)); err != nil {
				return err
			}
		}

		{
			var (
				name         = fmt.Sprintf("button-left-%d", buttonIndex)
				topLeftX     = 0
				topLeftY     = BoardMargin + (TileOuterSize * buttonIndex) + 1
				bottomRightX = topLeftX + TileSize
				bottomRightY = topLeftY + TileOuterSize
			)

			if err := g.drawButton(name, "\n→", topLeftX, topLeftY, bottomRightX, bottomRightY, g.insertTile(LeftButton, buttonIndex)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g gameLoop) drawRemainingTile(boardSize, boardOffset int) error {

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

		if err := g.drawButton(name, "\n ↶", topLeftX, topLeftY, bottomRightX, bottomRightY, g.rotateRemainingTile(RotateAntiClockWise)); err != nil {
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

		if err := g.drawButton(name, "\n ↷", topLeftX, topLeftY, bottomRightX, bottomRightY, g.rotateRemainingTile(RotateClockWise)); err != nil {
			return err
		}
	}

	return nil
}

func (g gameLoop) layout(gui *gocui.Gui) error {
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

func (g gameLoop) insertTile(button Button, buttonIndex int) GuiHandler {
	return func(gui *gocui.Gui, view *gocui.View) error {
		switch button {
		case TopButton:
			return g.state.InsertTileTopAt(buttonIndex)
		case RightButton:
			return g.state.InsertTileRightAt(buttonIndex)
		case BottomButton:
			return g.state.InsertTileBottomAt(buttonIndex)
		case LeftButton:
			return g.state.InsertTileLeftAt(buttonIndex)
		}
		return nil
	}
}

func (g gameLoop) rotateRemainingTile(rotationType RotationType) GuiHandler {
	return func(gui *gocui.Gui, view *gocui.View) error {
		switch rotationType {
		case RotateClockWise:
			g.state.RotateRemainingTileClockwise()
		case RotateAntiClockWise:
			g.state.RotateRemainingTileAntiClockwise()
		}
		return nil
	}
}

func (g gameLoop) quit(gui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (g *gameLoop) Run() error {
	gui, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		return err
	}
	defer gui.Close()
	g.gui = gui

	gui.SetManagerFunc(g.layout)

	gui.Cursor = true
	gui.Mouse = true

	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, g.quit); err != nil {
		return errors.Join(errors.New("failed to set exit key"), err)
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return errors.Join(errors.New("failed to start main loop"), err)
	}

	return nil
}

// GameLoop runs the labyrinth game with the provided initial state.
func GameLoop(initialState *model.Board) error {
	return (&gameLoop{
		boardDrawer: NewBoardDrawer(),
		state:       initialState,
		gui:         nil,
	}).Run()
}
