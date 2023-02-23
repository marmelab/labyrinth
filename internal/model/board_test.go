package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewTestBoard() *Board {
	return &Board{
		Tiles: [][]*BoardTile{
			{
				{
					Tile: &Tile{
						Shape:    ShapeI,
						Treasure: NoTreasure,
					},
					Rotation: Rotation0,
				}, {
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: 'B',
					},
					Rotation: Rotation270,
				}, {
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'A',
					},
					Rotation: Rotation180,
				},
			},
			{

				{
					Tile: &Tile{
						Shape:    ShapeI,
						Treasure: NoTreasure,
					},
					Rotation: Rotation90,
				}, {
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'C',
					},
					Rotation: Rotation270,
				}, {
					Tile: &Tile{
						Shape:    ShapeI,
						Treasure: NoTreasure,
					},
					Rotation: Rotation180,
				},
			},
			{

				{
					Tile: &Tile{
						Shape:    ShapeI,
						Treasure: NoTreasure,
					},
					Rotation: Rotation90,
				}, {
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'D',
					},
					Rotation: Rotation0,
				}, {
					Tile: &Tile{
						Shape:    ShapeI,
						Treasure: NoTreasure,
					},
					Rotation: Rotation180,
				},
			},
		},
		RemainingTile: &BoardTile{
			Tile: &Tile{
				Shape:    ShapeT,
				Treasure: 'E',
			},
			Rotation: Rotation180,
		},
		Players: []*Player{
			{
				Color: ColorBlue,
				Position: &Coordinate{
					Line: 1,
					Row:  1,
				},
				Targets: []Treasure{
					'B',
					'D',
				},
				Score: 0,
			},
			{
				Color: ColorGreen,
				Position: &Coordinate{
					Line: 1,
					Row:  1,
				},
				Targets: []Treasure{
					'C',
					'A',
				},
				Score: 0,
			},
		},
		RemainingPlayers:     []int{0, 1},
		RemainingPlayerIndex: 0,
	}
}

func TestCoordinates(t *testing.T) {
	t.Run("Contains()", func(t *testing.T) {
		t.Run("Should return true if the target coordinate is present in the array", func(t *testing.T) {
			coordinates := Coordinates{
				&Coordinate{0, 0},
				&Coordinate{0, 1},
			}

			assert.True(t, coordinates.Contains(&Coordinate{0, 0}))
		})

		t.Run("Should return false if the target coordinate is not present in the array", func(t *testing.T) {
			coordinates := Coordinates{
				&Coordinate{0, 0},
				&Coordinate{0, 1},
			}

			assert.False(t, coordinates.Contains(&Coordinate{1, 1}))
		})

	})
}

func TestBoard(t *testing.T) {

	t.Run("InsertTileTopAt()", func(t *testing.T) {
		t.Run("Should fail if row is not odd", func(t *testing.T) {
			err := NewTestBoard().InsertTileTopAt(0)
			assert.NotNil(t, err)
			assert.Equal(t, ErrEvenRow, err)
		})

		t.Run("Should fail if state is not in move tile", func(t *testing.T) {
			board := &Board{State: GameStateMovePawn}
			err := board.InsertTileTopAt(1)
			assert.NotNil(t, err)
			assert.Equal(t, ErrInvalidAction, err)
		})

		t.Run("Should slide all tiles bottom at inserted row", func(t *testing.T) {
			board := NewTestBoard()

			row := 1

			err := board.InsertTileTopAt(row)
			assert.Nil(t, err)

			assert.Equal(t, ShapeV, board.Tiles[1][row].Tile.Shape)
			assert.Equal(t, Treasure('B'), board.Tiles[1][row].Tile.Treasure)
			assert.Equal(t, Rotation270, board.Tiles[1][row].Rotation)

			assert.Equal(t, ShapeT, board.Tiles[2][row].Tile.Shape)
			assert.Equal(t, Treasure('C'), board.Tiles[2][row].Tile.Treasure)
			assert.Equal(t, Rotation270, board.Tiles[2][row].Rotation)
		})

		t.Run("Should insert remaining tile in first line for the given index", func(t *testing.T) {
			board := NewTestBoard()

			row := 1
			err := board.InsertTileTopAt(row)
			assert.Nil(t, err)

			assert.Equal(t, ShapeT, board.Tiles[0][row].Tile.Shape)
			assert.Equal(t, Treasure('E'), board.Tiles[0][row].Tile.Treasure)
			assert.Equal(t, Rotation180, board.Tiles[0][row].Rotation)
		})

		t.Run("Should set remaining tile t the last tile in the row", func(t *testing.T) {
			board := NewTestBoard()

			err := board.InsertTileTopAt(1)
			assert.Nil(t, err)

			assert.Equal(t, ShapeT, board.RemainingTile.Tile.Shape)
			assert.Equal(t, Treasure('D'), board.RemainingTile.Tile.Treasure)
			assert.Equal(t, Rotation0, board.RemainingTile.Rotation)
		})

		t.Run("Should set game state to MovePawn", func(t *testing.T) {
			board := NewTestBoard()

			err := board.InsertTileTopAt(1)
			assert.Nil(t, err)
			assert.Equal(t, GameStateMovePawn, board.State)
		})

		t.Run("Should move player if on the line", func(t *testing.T) {
			board := NewTestBoard()

			{
				err := board.InsertTileTopAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 2, board.Players[0].Position.Line)
				assert.Equal(t, 1, board.Players[0].Position.Row)
			}

			{
				board.State = GameStatePlaceTile
				err := board.InsertTileTopAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 0, board.Players[0].Position.Line)
				assert.Equal(t, 1, board.Players[0].Position.Row)
			}
		})
	})

	t.Run("InsertTileRightAt()", func(t *testing.T) {
		t.Run("Should fail if row is not odd", func(t *testing.T) {
			err := NewTestBoard().InsertTileRightAt(0)
			assert.NotNil(t, err)
			assert.Equal(t, ErrEvenRow, err)
		})

		t.Run("Should fail if state is not in move tile", func(t *testing.T) {
			board := &Board{State: GameStateMovePawn}
			err := board.InsertTileRightAt(1)
			assert.NotNil(t, err)
			assert.Equal(t, ErrInvalidAction, err)
		})

		t.Run("Should slide all tiles bottom at inserted row", func(t *testing.T) {
			board := NewTestBoard()

			line := 1

			err := board.InsertTileRightAt(line)
			assert.Nil(t, err)

			assert.Equal(t, ShapeT, board.Tiles[line][0].Tile.Shape)
			assert.Equal(t, Treasure('C'), board.Tiles[line][0].Tile.Treasure)
			assert.Equal(t, Rotation270, board.Tiles[line][0].Rotation)

			assert.Equal(t, ShapeI, board.Tiles[line][1].Tile.Shape)
			assert.Equal(t, NoTreasure, board.Tiles[line][1].Tile.Treasure)
			assert.Equal(t, Rotation180, board.Tiles[line][1].Rotation)
		})

		t.Run("Should insert remaining tile in the last row for the given index", func(t *testing.T) {
			board := NewTestBoard()

			line := 1
			err := board.InsertTileRightAt(line)
			assert.Nil(t, err)

			assert.Equal(t, ShapeT, board.Tiles[line][2].Tile.Shape)
			assert.Equal(t, Treasure('E'), board.Tiles[line][2].Tile.Treasure)
			assert.Equal(t, Rotation180, board.Tiles[line][2].Rotation)
		})

		t.Run("Should set remaining tile to the first tile in the row", func(t *testing.T) {
			board := NewTestBoard()

			err := board.InsertTileRightAt(1)
			assert.Nil(t, err)

			assert.Equal(t, ShapeI, board.RemainingTile.Tile.Shape)
			assert.Equal(t, NoTreasure, board.RemainingTile.Tile.Treasure)
			assert.Equal(t, Rotation90, board.RemainingTile.Rotation)
		})

		t.Run("Should set game state to MovePawn", func(t *testing.T) {
			board := NewTestBoard()

			err := board.InsertTileRightAt(1)
			assert.Nil(t, err)
			assert.Equal(t, GameStateMovePawn, board.State)
		})

		t.Run("Should move player if on the row", func(t *testing.T) {
			board := NewTestBoard()

			{
				err := board.InsertTileRightAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 1, board.Players[0].Position.Line)
				assert.Equal(t, 0, board.Players[0].Position.Row)
			}

			{
				board.State = GameStatePlaceTile
				err := board.InsertTileRightAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 1, board.Players[0].Position.Line)
				assert.Equal(t, 2, board.Players[0].Position.Row)
			}
		})
	})

	t.Run("InsertTileBottomAt()", func(t *testing.T) {
		t.Run("Should fail if row is not odd", func(t *testing.T) {
			err := NewTestBoard().InsertTileBottomAt(0)
			assert.NotNil(t, err)
			assert.Equal(t, ErrEvenRow, err)
		})

		t.Run("Should fail if state is not in move tile", func(t *testing.T) {
			board := &Board{State: GameStateMovePawn}
			err := board.InsertTileBottomAt(1)
			assert.NotNil(t, err)
			assert.Equal(t, ErrInvalidAction, err)
		})

		t.Run("Should slide all tiles bottom at inserted row", func(t *testing.T) {
			board := NewTestBoard()

			row := 1

			err := board.InsertTileBottomAt(row)
			assert.Nil(t, err)

			assert.Equal(t, ShapeT, board.Tiles[0][row].Tile.Shape)
			assert.Equal(t, Treasure('C'), board.Tiles[0][row].Tile.Treasure)
			assert.Equal(t, Rotation270, board.Tiles[0][row].Rotation)

			assert.Equal(t, ShapeT, board.Tiles[1][row].Tile.Shape)
			assert.Equal(t, Treasure('D'), board.Tiles[1][row].Tile.Treasure)
			assert.Equal(t, Rotation0, board.Tiles[1][row].Rotation)

		})

		t.Run("Should insert remaining tile in first line for the given index", func(t *testing.T) {
			board := NewTestBoard()

			row := 1
			err := board.InsertTileBottomAt(row)
			assert.Nil(t, err)

			assert.Equal(t, ShapeT, board.Tiles[2][row].Tile.Shape)
			assert.Equal(t, Treasure('E'), board.Tiles[2][row].Tile.Treasure)
			assert.Equal(t, Rotation180, board.Tiles[2][row].Rotation)
		})

		t.Run("Should set remaining tile t the last tile in the row", func(t *testing.T) {
			board := NewTestBoard()

			err := board.InsertTileBottomAt(1)
			assert.Nil(t, err)

			assert.Equal(t, ShapeV, board.RemainingTile.Tile.Shape)
			assert.Equal(t, Treasure('B'), board.RemainingTile.Tile.Treasure)
			assert.Equal(t, Rotation270, board.RemainingTile.Rotation)
		})

		t.Run("Should set game state to MovePawn", func(t *testing.T) {
			board := NewTestBoard()

			err := board.InsertTileBottomAt(1)
			assert.Nil(t, err)
			assert.Equal(t, GameStateMovePawn, board.State)
		})

		t.Run("Should move player if on the line", func(t *testing.T) {
			board := NewTestBoard()

			{
				err := board.InsertTileBottomAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 0, board.Players[0].Position.Line)
				assert.Equal(t, 1, board.Players[0].Position.Row)
			}

			{
				board.State = GameStatePlaceTile
				err := board.InsertTileBottomAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 1, board.Players[0].Position.Row)
				assert.Equal(t, 2, board.Players[0].Position.Line)
			}
		})
	})

	t.Run("InsertTileLeftAt()", func(t *testing.T) {
		t.Run("Should fail if row is not odd", func(t *testing.T) {
			err := NewTestBoard().InsertTileLeftAt(0)
			assert.NotNil(t, err)
			assert.Equal(t, ErrEvenRow, err)
		})

		t.Run("Should fail if state is not in move tile", func(t *testing.T) {
			board := &Board{State: GameStateMovePawn}
			err := board.InsertTileLeftAt(1)
			assert.NotNil(t, err)
			assert.Equal(t, ErrInvalidAction, err)
		})

		t.Run("Should slide all tiles bottom at inserted row", func(t *testing.T) {
			board := NewTestBoard()

			line := 1

			err := board.InsertTileLeftAt(line)
			assert.Nil(t, err)

			assert.Equal(t, ShapeI, board.Tiles[line][1].Tile.Shape)
			assert.Equal(t, NoTreasure, board.Tiles[line][1].Tile.Treasure)
			assert.Equal(t, Rotation90, board.Tiles[line][1].Rotation)

			assert.Equal(t, ShapeT, board.Tiles[line][2].Tile.Shape)
			assert.Equal(t, Treasure('C'), board.Tiles[line][2].Tile.Treasure)
			assert.Equal(t, Rotation270, board.Tiles[line][2].Rotation)
		})

		t.Run("Should insert remaining tile in the last row for the given index", func(t *testing.T) {
			board := NewTestBoard()

			line := 1
			err := board.InsertTileLeftAt(line)
			assert.Nil(t, err)

			assert.Equal(t, ShapeT, board.Tiles[line][0].Tile.Shape)
			assert.Equal(t, Treasure('E'), board.Tiles[line][0].Tile.Treasure)
			assert.Equal(t, Rotation180, board.Tiles[line][0].Rotation)
		})

		t.Run("Should set remaining tile to the first tile in the row", func(t *testing.T) {
			board := NewTestBoard()

			err := board.InsertTileLeftAt(1)
			assert.Nil(t, err)

			assert.Equal(t, ShapeI, board.RemainingTile.Tile.Shape)
			assert.Equal(t, NoTreasure, board.RemainingTile.Tile.Treasure)
			assert.Equal(t, Rotation180, board.RemainingTile.Rotation)
		})

		t.Run("Should set game state to MovePawn", func(t *testing.T) {
			board := NewTestBoard()

			err := board.InsertTileLeftAt(1)
			assert.Nil(t, err)
			assert.Equal(t, GameStateMovePawn, board.State)
		})

		t.Run("Should move player if on the row", func(t *testing.T) {
			board := NewTestBoard()

			{
				err := board.InsertTileLeftAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 1, board.Players[0].Position.Line)
				assert.Equal(t, 2, board.Players[0].Position.Row)
			}

			{
				board.State = GameStatePlaceTile
				err := board.InsertTileLeftAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 1, board.Players[0].Position.Line)
				assert.Equal(t, 0, board.Players[0].Position.Row)
			}
		})
	})

	t.Run("RotateRemainingTileClockwise()", func(t *testing.T) {
		t.Run("Should rotate remaining tile clockwise", func(t *testing.T) {
			board := &Board{
				RemainingTile: &BoardTile{
					Rotation: Rotation0,
				},
			}

			board.RotateRemainingTileClockwise()
			assert.Equal(t, Rotation90, board.RemainingTile.Rotation)

			board.RotateRemainingTileClockwise()
			assert.Equal(t, Rotation180, board.RemainingTile.Rotation)

			board.RotateRemainingTileClockwise()
			assert.Equal(t, Rotation270, board.RemainingTile.Rotation)

			board.RotateRemainingTileClockwise()
			assert.Equal(t, Rotation0, board.RemainingTile.Rotation)
		})
	})

	t.Run("RotateRemainingTileAntiClockwise()", func(t *testing.T) {
		t.Run("Should rotate remaining tile anticlockwise", func(t *testing.T) {
			board := &Board{
				RemainingTile: &BoardTile{
					Rotation: Rotation0,
				},
			}

			board.RotateRemainingTileAntiClockwise()
			assert.Equal(t, Rotation270, board.RemainingTile.Rotation)

			board.RotateRemainingTileAntiClockwise()
			assert.Equal(t, Rotation180, board.RemainingTile.Rotation)

			board.RotateRemainingTileAntiClockwise()
			assert.Equal(t, Rotation90, board.RemainingTile.Rotation)

			board.RotateRemainingTileAntiClockwise()
			assert.Equal(t, Rotation0, board.RemainingTile.Rotation)
		})
	})

	t.Run("MoveCurrentPlayerTo()", func(t *testing.T) {
		t.Run("Should return an error if state is not move player", func(t *testing.T) {
			board := NewTestBoard()

			err := board.MoveCurrentPlayerTo(1, 1)
			assert.Equal(t, ErrInvalidAction, err)
		})

		t.Run("Should allow to move on the same tile", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn
			board.GetCurrentPlayer().Targets = []Treasure{'E'}

			err := board.MoveCurrentPlayerTo(1, 1)
			assert.Nil(t, err)
			assert.Equal(t, 1, board.Players[0].Position.Line)
			assert.Equal(t, 1, board.Players[0].Position.Row)
		})

		t.Run("Should return an error if line is not valid", func(t *testing.T) {
			board := &Board{
				Tiles: make([][]*BoardTile, 3),
				Players: []*Player{
					{
						Color: ColorBlue,
						Position: &Coordinate{
							Line: 0,
							Row:  0,
						},
					},
				},
				State: GameStateMovePawn,
			}

			err := board.MoveCurrentPlayerTo(4, 1)
			assert.Equal(t, ErrInvalidAction, err)
		})

		t.Run("Should return an error if row is not valid", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn

			err := board.MoveCurrentPlayerTo(1, 4)
			assert.Equal(t, ErrInvalidAction, err)
		})

		t.Run("Should set player position", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn
			board.GetCurrentPlayer().Targets = []Treasure{'E'}

			err := board.MoveCurrentPlayerTo(0, 2)
			assert.Nil(t, err)
			assert.Equal(t, 0, board.Players[0].Position.Line)
			assert.Equal(t, 2, board.Players[0].Position.Row)
		})

		t.Run("Should set game state", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn
			board.GetCurrentPlayer().Targets = []Treasure{'E'}

			err := board.MoveCurrentPlayerTo(1, 1)
			assert.Nil(t, err)
			assert.Equal(t, GameStatePlaceTile, board.State)
		})

		t.Run("Should not increase player score if not on target treasure", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn
			board.GetCurrentPlayer().Targets = []Treasure{'E'}

			err := board.MoveCurrentPlayerTo(1, 1)
			assert.Nil(t, err)
			assert.Equal(t, 0, board.Players[0].Score)
		})

		t.Run("Should increase player score if on target treasure", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn

			err := board.MoveCurrentPlayerTo(0, 1)
			assert.Nil(t, err)
			assert.Equal(t, 1, board.Players[0].Score)
		})

		t.Run("Should not pop treasure from hand if not on target treasure", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn
			board.GetCurrentPlayer().Targets = []Treasure{'E'}

			err := board.MoveCurrentPlayerTo(1, 1)
			assert.Nil(t, err)
			assert.Equal(t, Treasure('E'), board.Players[0].Targets[0])
		})

		t.Run("Should pop treasure from hand if on target treasure", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn

			err := board.MoveCurrentPlayerTo(0, 1)
			assert.Nil(t, err)
			assert.Equal(t, Treasure('D'), board.Players[0].Targets[0])
		})

		t.Run("Should not remove treasure from tile if not on target treasure", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn
			board.GetCurrentPlayer().Targets = []Treasure{'E'}

			err := board.MoveCurrentPlayerTo(1, 1)
			assert.Nil(t, err)
			assert.Equal(t, Treasure('E'), board.RemainingTile.Tile.Treasure)
		})

		t.Run("Should remove treasure form tile if on target treasure", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn

			err := board.MoveCurrentPlayerTo(0, 1)
			assert.Nil(t, err)
			assert.Equal(t, NoTreasure, board.Tiles[0][1].Tile.Treasure)
		})

		t.Run("Should set to game end if only one player remains", func(t *testing.T) {
			board := NewTestBoard()
			board.RemainingPlayers = []int{0}
			board.Players[0].Targets = []Treasure{'B'}
			{
				board.State = GameStateMovePawn
				err := board.MoveCurrentPlayerTo(0, 1)
				assert.Nil(t, err)
				assert.Equal(t, GameStateEnd, board.State)
			}
		})

		t.Run("Should increase current player index", func(t *testing.T) {
			board := NewTestBoard()

			{
				// Blue
				board.State = GameStateMovePawn
				board.GetCurrentPlayer().Targets = []Treasure{'E'}
				err := board.MoveCurrentPlayerTo(0, 1)
				assert.Nil(t, err)
				assert.Equal(t, 1, board.RemainingPlayerIndex)
			}

			{
				// Green
				board.State = GameStateMovePawn
				board.GetCurrentPlayer().Targets = []Treasure{'E'}
				err := board.MoveCurrentPlayerTo(1, 1)
				assert.Nil(t, err)
				assert.Equal(t, 0, board.RemainingPlayerIndex)
			}
		})

		t.Run("Should drop player from remaining player if he does not have any more treasures to fetch", func(t *testing.T) {

			{
				// Blue
				board := NewTestBoard()
				board.State = GameStateMovePawn

				board.Players[0].Targets = []Treasure{'D'}
				err := board.MoveCurrentPlayerTo(2, 1)
				assert.Nil(t, err)
				assert.Equal(t, 0, board.RemainingPlayerIndex)
				assert.Equal(t, 1, len(board.RemainingPlayers))
				assert.Equal(t, 1, board.RemainingPlayers[0])
			}

			{
				// Green
				board := NewTestBoard()
				board.State = GameStateMovePawn

				board.Players[1].Targets = []Treasure{'A'}
				board.RemainingPlayerIndex = 1
				err := board.MoveCurrentPlayerTo(0, 2)
				assert.Nil(t, err)
				assert.Equal(t, 0, board.RemainingPlayerIndex)
				assert.Equal(t, 1, len(board.RemainingPlayers))
				assert.Equal(t, 0, board.RemainingPlayers[0])
			}
		})
	})

	t.Run("GetCurrentPlayer()", func(t *testing.T) {
		t.Run("Should return the current player", func(t *testing.T) {
			bluePlayer := &Player{
				Color: ColorBlue,
				Position: &Coordinate{
					Line: 0,
					Row:  0,
				},
			}

			greenPlayer := &Player{
				Color: ColorGreen,
				Position: &Coordinate{
					Line: 0,
					Row:  0,
				},
			}

			board := &Board{
				Players:              []*Player{bluePlayer, greenPlayer},
				RemainingPlayers:     []int{0, 1},
				RemainingPlayerIndex: 0,
			}
			assert.Equal(t, bluePlayer, board.GetCurrentPlayer())

			board.RemainingPlayerIndex = 1
			assert.Equal(t, greenPlayer, board.GetCurrentPlayer())
		})
	})

	t.Run("getAccessibleNeighbors()", func(t *testing.T) {
		t.Run("Should return all accessible neighbors for a tile", func(t *testing.T) {
			tests := []struct {
				line     int
				row      int
				expected Coordinates
			}{
				{0, 0, Coordinates{}},
				{0, 2, Coordinates{
					{0, 1},
				}},
				{2, 0, Coordinates{
					{1, 0},
				}},
				{1, 0, Coordinates{
					{2, 0},
				}},
				{2, 1, Coordinates{
					{1, 1},
					{2, 2},
				}},
			}

			board := NewTestBoard()
			for _, test := range tests {
				neighbors := board.getAccessibleNeighbors(test.line, test.row)
				assert.Equal(t, test.expected, neighbors)
			}
		})
	})

	t.Run("getAccessibleTilesForCoordinate()", func(t *testing.T) {
		t.Run("Should return all accessible tiles from given coordinate", func(t *testing.T) {
			board := NewTestBoard()

			tiles := board.getAccessibleTilesForCoordinate(&Coordinate{0, 2})
			assert.Equal(t, 5, len(tiles))
			assert.Equal(t, 0, tiles[0].Line)
			assert.Equal(t, 2, tiles[0].Row)
			assert.Equal(t, 0, tiles[1].Line)
			assert.Equal(t, 1, tiles[1].Row)
			assert.Equal(t, 1, tiles[2].Line)
			assert.Equal(t, 1, tiles[2].Row)
			assert.Equal(t, 2, tiles[3].Line)
			assert.Equal(t, 1, tiles[3].Row)
			assert.Equal(t, 2, tiles[4].Line)
			assert.Equal(t, 2, tiles[4].Row)
		})
	})

	t.Run("GetAccessibleTiles()", func(t *testing.T) {
		t.Run("Should return all accessible tiles from given player", func(t *testing.T) {
			board := NewTestBoard()
			board.Players[0].Targets = []Treasure{'E'}

			assert.Equal(t, Coordinates{
				{1, 1},
				{0, 1},
				{2, 1},
				{0, 2},
				{2, 2},
			}, board.GetAccessibleTiles())
		})
	})

	t.Run("GetSize()", func(t *testing.T) {
		t.Run("Should return board size", func(t *testing.T) {
			board := &Board{
				Tiles: make([][]*BoardTile, 3),
			}
			assert.Equal(t, 3, board.GetSize())
		})
	})

	t.Run("getShortestPath()", func(t *testing.T) {
		t.Run("Should return the shortest path between current player and its target", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn
			board.Players[0].Position.Line = 2
			board.Players[0].Position.Row = 1

			path := board.getShortestPath()
			assert.Equal(t, Coordinates{
				{1, 1},
				{0, 1},
			}, path)
		})
	})
}
func TestNewBoard(t *testing.T) {
	t.Run("Should return an error if size is even.", func(t *testing.T) {
		board, err := NewBoard(2, 1)
		assert.NotNil(t, err)
		assert.Equal(t, "the board size must be either 3 or 7, got: 2", err.Error())
		assert.Nil(t, board)
	})

	t.Run("Should return an error if player count is not between 1 and 4.", func(t *testing.T) {
		{
			_, err := NewBoard(3, 0)
			assert.NotNil(t, err)
			assert.Equal(t, "the number of players must be between 1 and 4 included, got: 0", err.Error())
		}

		for playerCount := 1; playerCount <= 4; playerCount++ {
			_, err := NewBoard(3, playerCount)
			assert.Nil(t, err)
		}

		{
			_, err := NewBoard(3, 5)
			assert.NotNil(t, err)
			assert.Equal(t, "the number of players must be between 1 and 4 included, got: 5", err.Error())
		}
	})

	t.Run("Should return a board instance if size is odd", func(t *testing.T) {
		board, err := NewBoard(3, 1)
		assert.Nil(t, err)
		assert.NotNil(t, board)
	})

	t.Run("Should initialize tiles in board", func(t *testing.T) {
		{
			size := 3
			board, _ := NewBoard(size, 1)
			assert.Equal(t, size, len(board.Tiles))
			for i := 0; i < size; i++ {
				assert.Equal(t, size, len(board.Tiles[i]))
				for j := 0; j < size; j++ {
					boardTile := board.Tiles[i][j]
					assert.NotNil(t, boardTile)
					assert.NotNil(t, boardTile.Tile)
				}
			}
		}
	})

	t.Run("Should initialize players", func(t *testing.T) {
		var expectedPlayers = []*Player{
			{
				Color: ColorBlue,
				Position: &Coordinate{
					Line: 0,
					Row:  0,
				},
			},
			{
				Color: ColorGreen,
				Position: &Coordinate{
					Line: 2,
					Row:  2,
				},
			},
			{
				Color: ColorRed,
				Position: &Coordinate{
					Line: 0,
					Row:  2,
				},
			},
			{
				Color: ColorYellow,
				Position: &Coordinate{
					Line: 2,
					Row:  0,
				},
			},
		}

		for playerCount := 1; playerCount <= 4; playerCount++ {

			board, _ := NewBoard(3, playerCount)

			assert.Equal(t, playerCount, len(board.Players))
			for i := 0; i < playerCount; i++ {
				assert.Equal(t, expectedPlayers[i].Color, board.Players[i].Color)
				assert.Equal(t, expectedPlayers[i].Position, board.Players[i].Position)
			}
		}
	})

	t.Run("Should initialize player targets", func(t *testing.T) {

		tests := [][]int{
			{3, 1, 5},
			{7, 1, 24},
			{3, 2, 2},
			{7, 2, 12},
			{3, 3, 1},
			{7, 3, 8},
			{3, 4, 1},
			{7, 4, 6},
		}

		for _, test := range tests {
			board, _ := NewBoard(test[0], test[1])
			for i := 0; i < test[1]; i++ {
				assert.Equal(t, test[2], len(board.Players[i].Targets))
			}
		}
	})

	t.Run("Should initialize remaining tile", func(t *testing.T) {
		{
			board, _ := NewBoard(3, 1)
			assert.NotNil(t, board.RemainingTile)
			assert.NotNil(t, board.RemainingTile.Tile)
			assert.Equal(t, Rotation0, board.RemainingTile.Rotation)
		}
	})
}

func TestGenerateTiles(t *testing.T) {
	t.Run("Should generate n**2 -3 tiles", func(t *testing.T) {
		tiles, _ := generateTiles(3)
		assert.Equal(t, 6, len(tiles))
	})

	t.Run("Should generate about 36% of T shaped tiles", func(t *testing.T) {
		{
			tiles, _ := generateTiles(3)
			for i := 0; i < 4; i++ {
				assert.Equal(t, ShapeT, tiles[i].Shape)
			}
			assert.Equal(t, Treasure('A'), tiles[0].Treasure)
			assert.Equal(t, Treasure('B'), tiles[1].Treasure)
			assert.Equal(t, Treasure('C'), tiles[2].Treasure)
			assert.Equal(t, Treasure('D'), tiles[3].Treasure)
		}
		{
			tiles, _ := generateTiles(7)
			for i := 0; i < 18; i++ {
				assert.Equal(t, ShapeT, tiles[i].Shape)
			}
		}
	})

	t.Run("Should generate about 11% of V shaped tiles with treasure", func(t *testing.T) {
		{
			tiles, _ := generateTiles(3)
			for i := 4; i < 5; i++ {
				assert.Equal(t, ShapeV, tiles[i].Shape)
			}
			assert.Equal(t, Treasure('E'), tiles[4].Treasure)
		}
		{
			tiles, _ := generateTiles(7)
			for i := 18; i < 24; i++ {
				assert.Equal(t, ShapeV, tiles[i].Shape)
			}
		}
	})

	t.Run("Should generate about 26% of I shaped tiles", func(t *testing.T) {
		{
			tiles, _ := generateTiles(7)
			for i := 24; i < 37; i++ {
				assert.Equal(t, ShapeI, tiles[i].Shape)
				assert.Equal(t, NoTreasure, tiles[i].Treasure)
			}
		}
	})

	t.Run("Should generate about 18% of V shaped tiles without treasure", func(t *testing.T) {
		{
			tiles, _ := generateTiles(7)
			for i := 37; i < 46; i++ {
				assert.Equal(t, ShapeV, tiles[i].Shape)
				assert.Equal(t, NoTreasure, tiles[i].Treasure)
			}
		}
	})
}
