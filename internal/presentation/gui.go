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

type gameLoop struct {
	boardDrawer BoardDrawer
	state       *model.Board
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
				topLeftX     = BoardMargin + (TileSize * buttonIndex)
				topLeftY     = BoardMargin + (boardSize + 1)
				bottomRightX = topLeftX + TileSize
				bottomRightY = topLeftY + TileSize
			)

			if bottomButton, err := gui.SetView(fmt.Sprintf("button-bottom-%d", buttonIndex), topLeftX, topLeftY, bottomRightX, bottomRightY); err != nil {
				if err != gocui.ErrUnknownView {
					return errors.Join(fmt.Errorf("failed to initialize left button at row %d", buttonIndex), err)
				}
				fmt.Fprint(bottomButton, "↑")
			}
		}

		{
			var (
				topLeftX     = 0
				topLeftY     = BoardMargin + (TileSize * buttonIndex)
				bottomRightX = topLeftX + TileSize
				bottomRightY = topLeftY + TileSize
			)

			if leftButton, err := gui.SetView(fmt.Sprintf("button-left-%d", buttonIndex), topLeftX, topLeftY, bottomRightX, bottomRightY); err != nil {
				if err != gocui.ErrUnknownView {
					return errors.Join(fmt.Errorf("failed to initialize left button at row %d", buttonIndex), err)
				}
				fmt.Fprint(leftButton, "→")
			}
		}
	}

	return nil
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
