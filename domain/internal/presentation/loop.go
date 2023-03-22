package presentation

import (
	"errors"

	"github.com/awesome-gocui/gocui"
	"github.com/marmelab/labyrinth/domain/internal/model"
	"github.com/marmelab/labyrinth/domain/internal/storage"
)

const (
	BoardMargin = 5
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

	board     *model.Board
	saveBoard storage.BoardSaveFn

	gui *gocui.Gui
}

func (g gameLoop) insertTile(button Button, buttonIndex int) GuiHandler {
	return func(gui *gocui.Gui, view *gocui.View) error {
		// We need to swallow this error case here, otherwise the program
		// will stop due to the library.
		// TODO: provide a clear error message to the UI.
		if g.board.State != model.GameStatePlaceTile {
			return nil
		}

		switch button {
		case TopButton:
			g.board.InsertTileTopAt(buttonIndex)
		case RightButton:
			g.board.InsertTileRightAt(buttonIndex)
		case BottomButton:
			g.board.InsertTileBottomAt(buttonIndex)
		case LeftButton:
			g.board.InsertTileLeftAt(buttonIndex)
		}

		return g.saveBoard()
	}
}

func (g gameLoop) rotateRemainingTile(rotationType RotationType) GuiHandler {
	return func(gui *gocui.Gui, view *gocui.View) error {
		switch rotationType {
		case RotateClockWise:
			g.board.RotateRemainingTileClockwise()
		case RotateAntiClockWise:
			g.board.RotateRemainingTileAntiClockwise()
		}
		return g.saveBoard()
	}
}

func (g gameLoop) moveCurrentPlayerTo(line, row int) GuiHandler {
	return func(gui *gocui.Gui, view *gocui.View) error {
		if err := g.board.MoveCurrentPlayerTo(line, row); err != nil {
			// We need to swallow this error case here, otherwise the program
			// will stop due to the library.
			// TODO: provide a clear error message to the UI.
			if err != model.ErrInvalidAction {
				return err
			}
		}
		return g.saveBoard()
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

	gameUi := &gameUi{
		gui:  gui,
		loop: g,

		boardDrawer: g.boardDrawer,
		board:       g.board,
	}

	gui.SetManagerFunc(gameUi.layout)

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

// RunGameLoop runs the labyrinth game with the provided initial state.
func RunGameLoop(initialState *model.Board, saveBoard storage.BoardSaveFn) error {
	return (&gameLoop{
		boardDrawer: NewBoardDrawer(),

		board:     initialState,
		saveBoard: saveBoard,

		gui: nil,
	}).Run()
}
