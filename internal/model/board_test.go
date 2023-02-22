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
						Shape:    ShapeV,
						Treasure: 'C',
					},
					Rotation: Rotation90,
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
						Shape:    ShapeV,
						Treasure: 'D',
					},
					Rotation: Rotation90,
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
				Line:  1,
				Row:   1,
				Targets: []Treasure{
					'B',
					'E',
					'C',
					'D',
					'A',
				},
				Score: 0,
			},
		},
	}
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

			assert.Equal(t, ShapeV, board.Tiles[2][row].Tile.Shape)
			assert.Equal(t, Treasure('C'), board.Tiles[2][row].Tile.Treasure)
			assert.Equal(t, Rotation90, board.Tiles[2][row].Rotation)
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

			assert.Equal(t, ShapeV, board.RemainingTile.Tile.Shape)
			assert.Equal(t, Treasure('D'), board.RemainingTile.Tile.Treasure)
			assert.Equal(t, Rotation90, board.RemainingTile.Rotation)
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
				assert.Equal(t, 2, board.Players[0].Line)
				assert.Equal(t, 1, board.Players[0].Row)
			}

			{
				board.State = GameStatePlaceTile
				err := board.InsertTileTopAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 0, board.Players[0].Line)
				assert.Equal(t, 1, board.Players[0].Row)
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

			assert.Equal(t, ShapeV, board.Tiles[line][0].Tile.Shape)
			assert.Equal(t, Treasure('C'), board.Tiles[line][0].Tile.Treasure)
			assert.Equal(t, Rotation90, board.Tiles[line][0].Rotation)

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
				assert.Equal(t, 1, board.Players[0].Line)
				assert.Equal(t, 0, board.Players[0].Row)
			}

			{
				board.State = GameStatePlaceTile
				err := board.InsertTileRightAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 1, board.Players[0].Line)
				assert.Equal(t, 2, board.Players[0].Row)
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

			assert.Equal(t, ShapeV, board.Tiles[0][row].Tile.Shape)
			assert.Equal(t, Treasure('C'), board.Tiles[0][row].Tile.Treasure)
			assert.Equal(t, Rotation90, board.Tiles[0][row].Rotation)

			assert.Equal(t, ShapeV, board.Tiles[1][row].Tile.Shape)
			assert.Equal(t, Treasure('D'), board.Tiles[1][row].Tile.Treasure)
			assert.Equal(t, Rotation90, board.Tiles[1][row].Rotation)

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
				assert.Equal(t, 0, board.Players[0].Line)
				assert.Equal(t, 1, board.Players[0].Row)
			}

			{
				board.State = GameStatePlaceTile
				err := board.InsertTileBottomAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 2, board.Players[0].Line)
				assert.Equal(t, 1, board.Players[0].Row)
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

			assert.Equal(t, ShapeV, board.Tiles[line][2].Tile.Shape)
			assert.Equal(t, Treasure('C'), board.Tiles[line][2].Tile.Treasure)
			assert.Equal(t, Rotation90, board.Tiles[line][2].Rotation)

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
				assert.Equal(t, 1, board.Players[0].Line)
				assert.Equal(t, 2, board.Players[0].Row)
			}

			{
				board.State = GameStatePlaceTile
				err := board.InsertTileLeftAt(1)
				assert.Nil(t, err)
				assert.Equal(t, 1, board.Players[0].Line)
				assert.Equal(t, 0, board.Players[0].Row)
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

			err := board.MoveCurrentPlayerTo(0, 0)
			assert.Nil(t, err)
			assert.Equal(t, 0, board.Players[0].Line)
			assert.Equal(t, 0, board.Players[0].Row)
		})

		t.Run("Should return an error if line is not valid", func(t *testing.T) {
			board := &Board{
				Tiles: make([][]*BoardTile, 3),
				Players: []*Player{
					{
						Color: ColorBlue,
						Line:  0,
						Row:   0,
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

			err := board.MoveCurrentPlayerTo(1, 2)
			assert.Nil(t, err)
			assert.Equal(t, 1, board.Players[0].Line)
			assert.Equal(t, 2, board.Players[0].Row)
		})

		t.Run("Should set game state", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn

			err := board.MoveCurrentPlayerTo(1, 2)
			assert.Nil(t, err)
			assert.Equal(t, GameStatePlaceTile, board.State)
		})

		t.Run("Should not increase player score if not on target treasure", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn

			err := board.MoveCurrentPlayerTo(0, 0)
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

			err := board.MoveCurrentPlayerTo(0, 0)
			assert.Nil(t, err)
			assert.Equal(t, Treasure('B'), board.Players[0].Targets[0])
		})

		t.Run("Should pop treasure from hand if on target treasure", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn

			err := board.MoveCurrentPlayerTo(0, 1)
			assert.Nil(t, err)
			assert.Equal(t, Treasure('E'), board.Players[0].Targets[0])
		})

		t.Run("Should not remove treasure from tile if not on target treasure", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn

			err := board.MoveCurrentPlayerTo(0, 0)
			assert.Nil(t, err)
			assert.Equal(t, Treasure('B'), board.Tiles[0][1].Tile.Treasure)
		})

		t.Run("Should remove treasure form tile if on target treasure", func(t *testing.T) {
			board := NewTestBoard()
			board.State = GameStateMovePawn

			err := board.MoveCurrentPlayerTo(0, 1)
			assert.Nil(t, err)
			assert.Equal(t, NoTreasure, board.Tiles[0][1].Tile.Treasure)
		})

		t.Run("Should set to game end if player has an empty hand", func(t *testing.T) {
			board := NewTestBoard()
			board.Players[0].Hand = []Treasure{'B'}
			{
				board.State = GameStateMovePawn
				err := board.MoveCurrentPlayerTo(0, 1)
				assert.Nil(t, err)
				assert.Equal(t, GameStateEnd, board.State)
			}
		})
	})

	t.Run("CurrentPlayer()", func(t *testing.T) {
		t.Run("Should return the current player", func(t *testing.T) {
			bluePlayer := &Player{
				Color: ColorBlue,
				Line:  0,
				Row:   0,
			}
			board := &Board{
				Players: []*Player{bluePlayer},
			}
			assert.Equal(t, bluePlayer, board.CurrentPlayer())
		})
	})

	t.Run("Size()", func(t *testing.T) {
		t.Run("Should return board size", func(t *testing.T) {
			board := &Board{
				Tiles: make([][]*BoardTile, 3),
			}
			assert.Equal(t, 3, board.Size())
		})
	})
}

func TestNewBoard(t *testing.T) {
	t.Run("Should return an error if size is even.", func(t *testing.T) {
		board, err := NewBoard(2)
		assert.NotNil(t, err)
		assert.Equal(t, "size must be an odd number, got: 2", err.Error())
		assert.Nil(t, board)
	})

	t.Run("Should return a board instance if size is odd", func(t *testing.T) {
		board, err := NewBoard(3)
		assert.Nil(t, err)
		assert.NotNil(t, board)
	})

	t.Run("Should initialize tiles in board", func(t *testing.T) {
		{
			size := 3
			board, _ := NewBoard(size)
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
		{
			board, _ := NewBoard(3)
			assert.Equal(t, 1, len(board.Players))
			assert.Equal(t, ColorBlue, board.Players[0].Color)
			assert.Equal(t, 0, board.Players[0].Line)
			assert.Equal(t, 0, board.Players[0].Row)
		}
	})

	t.Run("Should initialize remaining tile", func(t *testing.T) {
		{
			board, _ := NewBoard(3)
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
