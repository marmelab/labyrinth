package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardTile(t *testing.T) {

	t.Run("Exits()", func(t *testing.T) {
		t.Run("Should return exists for ", func(t *testing.T) {
			tests := []struct {
				shape    Shape
				rotation Rotation
				exits    TileExits
			}{
				{ShapeI, Rotation0, TileExits{TileExitRight, TileExitLeft}},
				{ShapeI, Rotation90, TileExits{TileExitTop, TileExitBottom}},
				{ShapeI, Rotation180, TileExits{TileExitRight, TileExitLeft}},
				{ShapeI, Rotation270, TileExits{TileExitTop, TileExitBottom}},
				{ShapeT, Rotation0, TileExits{TileExitTop, TileExitRight, TileExitLeft}},
				{ShapeT, Rotation90, TileExits{TileExitTop, TileExitRight, TileExitBottom}},
				{ShapeT, Rotation180, TileExits{TileExitRight, TileExitBottom, TileExitLeft}},
				{ShapeT, Rotation270, TileExits{TileExitTop, TileExitBottom, TileExitLeft}},
				{ShapeV, Rotation0, TileExits{TileExitBottom, TileExitLeft}},
				{ShapeV, Rotation90, TileExits{TileExitTop, TileExitLeft}},
				{ShapeV, Rotation180, TileExits{TileExitTop, TileExitRight}},
				{ShapeV, Rotation270, TileExits{TileExitRight, TileExitBottom}},
			}

			for _, test := range tests {
				boardTile := &BoardTile{
					Tile: &Tile{
						Shape: test.shape,
					},
					Rotation: test.rotation,
				}
				assert.Equal(t, test.exits, boardTile.GetExits())
			}
		})
	})
}
