package presentation

import (
	"errors"
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/marmelab/labyrinth/internal/model"
)

const (
	BoardMargin = 4
)

type Button int

const (
	TopButton = iota
	RightButton
	BottomButton
	LeftButton
)

type GuiHandler func(gui *gocui.Gui, view *gocui.View) error

type gameLoop struct {
	boardDrawer BoardDrawer
	state       *model.Board
}

func (g gameLoop) button(gui *gocui.Gui, name, text string, topLeftX, topLeftY, bottomRightX, bottomRightY int, handler GuiHandler) (*gocui.View, error) {
	button, err := gui.SetView(name, topLeftX, topLeftY, bottomRightX, bottomRightY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return nil, errors.Join(fmt.Errorf("failed to initialize button %s", name), err)
		}

		button.Highlight = true
		button.SelBgColor = gocui.ColorGreen
		button.SelFgColor = gocui.ColorBlack

		fmt.Fprint(button, text)
		if err := gui.SetKeybinding(name, gocui.MouseLeft, gocui.ModNone, handler); err != nil {
			return nil, errors.Join(fmt.Errorf("failed to set mouse click to %s", name), err)
		}
	}

	return button, nil
}

func (g gameLoop) layout(gui *gocui.Gui) error {
	var (
		tileCount = len(g.state.Tiles)
		boardSize = tileCount * TileSize
	)

	if boardView, err := gui.SetView("board", BoardMargin, BoardMargin, boardSize+BoardMargin, boardSize+BoardMargin); err != nil {
		if err != gocui.ErrUnknownView {
			return errors.Join(errors.New("failed to initialize board"), err)
		}
		g.boardDrawer.DrawTo(boardView, g.state)
	}

	for buttonIndex := 1; buttonIndex < tileCount; buttonIndex += 2 {
		{
			var (
				name         = fmt.Sprintf("button-top-%d", buttonIndex)
				topLeftX     = BoardMargin + (TileSize * buttonIndex)
				topLeftY     = 0
				bottomRightX = topLeftX + TileSize
				bottomRightY = topLeftY + TileSize
			)

			g.button(gui, name, "↓ ", topLeftX, topLeftY, bottomRightX, bottomRightY, g.insertTile(TopButton, buttonIndex))
		}

		{
			var (
				name         = fmt.Sprintf("button-right-%d", buttonIndex)
				topLeftX     = BoardMargin + (boardSize + 1)
				topLeftY     = BoardMargin + (TileSize * buttonIndex)
				bottomRightX = topLeftX + TileSize
				bottomRightY = topLeftY + TileSize
			)

			g.button(gui, name, "← ", topLeftX, topLeftY, bottomRightX, bottomRightY, g.insertTile(RightButton, buttonIndex))
		}

		{
			var (
				name         = fmt.Sprintf("button-bottom-%d", buttonIndex)
				topLeftX     = BoardMargin + (TileSize * buttonIndex)
				topLeftY     = BoardMargin + (boardSize + 1)
				bottomRightX = topLeftX + TileSize
				bottomRightY = topLeftY + TileSize
			)

			g.button(gui, name, "↑ ", topLeftX, topLeftY, bottomRightX, bottomRightY, g.insertTile(BottomButton, buttonIndex))
		}

		{
			var (
				name         = fmt.Sprintf("button-left-%d", buttonIndex)
				topLeftX     = 0
				topLeftY     = BoardMargin + (TileSize * buttonIndex)
				bottomRightX = topLeftX + TileSize
				bottomRightY = topLeftY + TileSize
			)

			g.button(gui, name, "→ ", topLeftX, topLeftY, bottomRightX, bottomRightY, g.insertTile(LeftButton, buttonIndex))
		}
	}

	return nil
}

func (g gameLoop) insertTile(button Button, buttonIndex int) GuiHandler {
	return func(gui *gocui.Gui, view *gocui.View) error {
		fmt.Println(button, buttonIndex)
		return nil
	}
}

func (g gameLoop) quit(gui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (g gameLoop) Run() error {
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	defer gui.Close()

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
	}).Run()
}
