package presentation

import (
	"errors"

	"github.com/awesome-gocui/gocui"
	"github.com/marmelab/labyrinth/internal/model"
	"github.com/marmelab/labyrinth/internal/storage"
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

	state     *model.Board
	saveBoard storage.BoardSaveFn

	gui *gocui.Gui
}

func (g gameLoop) insertTile(button Button, buttonIndex int) GuiHandler {
	return func(gui *gocui.Gui, view *gocui.View) error {
		switch button {
		case TopButton:
			if err := g.state.InsertTileTopAt(buttonIndex); err != nil {
				return err
			}
		case RightButton:
			if err := g.state.InsertTileRightAt(buttonIndex); err != nil {
				return err
			}
		case BottomButton:
			if err := g.state.InsertTileBottomAt(buttonIndex); err != nil {
				return err
			}
		case LeftButton:
			if err := g.state.InsertTileLeftAt(buttonIndex); err != nil {
				return err
			}
		}

		return g.saveBoard()
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
		state:       g.state,
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
		state:       initialState,
		saveBoard:   saveBoard,
		gui:         nil,
	}).Run()
}
