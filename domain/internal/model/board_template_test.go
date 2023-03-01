package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
