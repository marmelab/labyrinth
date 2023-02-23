package presentation

import (
	"strings"
	"testing"

	"github.com/marmelab/labyrinth/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestBoardDrawer(t *testing.T) {

	t.Run("Should draw a tile", func(t *testing.T) {
		drawer := &boardDrawer{}

		// I-shaped til
		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeI,
					Treasure: model.NoTreasure,
				},
				Rotation: model.Rotation0,
			})

			assert.Equal(t, `
───
 · 
───
`[1:], writer.String())
		}

		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeI,
					Treasure: model.NoTreasure,
				},
				Rotation: model.Rotation90,
			})

			assert.Equal(t, `
│ │
│·│
│ │
`[1:], writer.String())
		}

		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeI,
					Treasure: model.NoTreasure,
				},
				Rotation: model.Rotation180,
			})

			assert.Equal(t, `
───
 · 
───
`[1:], writer.String())
		}

		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeI,
					Treasure: model.NoTreasure,
				},
				Rotation: model.Rotation270,
			})

			assert.Equal(t, `
│ │
│·│
│ │
`[1:], writer.String())
		}

		// T-shaped til
		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeT,
					Treasure: 'A',
				},
				Rotation: model.Rotation0,
			})

			assert.Equal(t, `
┘ └
 A 
───
`[1:], writer.String())
		}

		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeT,
					Treasure: 'B',
				},
				Rotation: model.Rotation90,
			})

			assert.Equal(t, `
│ └
│B 
│ ┌
`[1:], writer.String())
		}

		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeT,
					Treasure: 'C',
				},
				Rotation: model.Rotation180,
			})

			assert.Equal(t, `
───
 C 
┐ ┌
`[1:], writer.String())
		}

		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeT,
					Treasure: 'D',
				},
				Rotation: model.Rotation270,
			})

			assert.Equal(t, `
┘ │
 D│
┐ │
`[1:], writer.String())
		}

		// V-shaped til
		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeV,
					Treasure: 'A',
				},
				Rotation: model.Rotation0,
			})

			assert.Equal(t, `
──┐
 A│
┐ │
`[1:], writer.String())
		}

		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeV,
					Treasure: 'B',
				},
				Rotation: model.Rotation90,
			})

			assert.Equal(t, `
┘ │
 B│
──┘
`[1:], writer.String())
		}

		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeV,
					Treasure: model.NoTreasure,
				},
				Rotation: model.Rotation180,
			})

			assert.Equal(t, `
│ └
│· 
└──
`[1:], writer.String())
		}

		{
			writer := new(strings.Builder)
			drawer.DrawBoardTileTo(writer, &model.BoardTile{
				Tile: &model.Tile{
					Shape:    model.ShapeV,
					Treasure: model.NoTreasure,
				},
				Rotation: model.Rotation270,
			})

			assert.Equal(t, `
┌──
│· 
│ ┌
`[1:], writer.String())
		}
	})

	t.Run("Should draw a board", func(t *testing.T) {
		drawer := &boardDrawer{}

		writer := new(strings.Builder)
		drawer.DrawTo(writer, &model.Board{
			Tiles: [][]*model.BoardTile{
				{
					{
						Tile: &model.Tile{
							Shape:    model.ShapeI,
							Treasure: model.NoTreasure,
						},
						Rotation: model.Rotation0,
					}, {
						Tile: &model.Tile{
							Shape:    model.ShapeV,
							Treasure: 'B',
						},
						Rotation: model.Rotation270,
					}, {
						Tile: &model.Tile{
							Shape:    model.ShapeT,
							Treasure: 'A',
						},
						Rotation: model.Rotation180,
					},
				},
				{

					{
						Tile: &model.Tile{
							Shape:    model.ShapeI,
							Treasure: model.NoTreasure,
						},
						Rotation: model.Rotation90,
					}, {
						Tile: &model.Tile{
							Shape:    model.ShapeV,
							Treasure: 'C',
						},
						Rotation: model.Rotation90,
					}, {
						Tile: &model.Tile{
							Shape:    model.ShapeI,
							Treasure: model.NoTreasure,
						},
						Rotation: model.Rotation180,
					},
				},
				{

					{
						Tile: &model.Tile{
							Shape:    model.ShapeI,
							Treasure: model.NoTreasure,
						},
						Rotation: model.Rotation90,
					}, {
						Tile: &model.Tile{
							Shape:    model.ShapeV,
							Treasure: 'D',
						},
						Rotation: model.Rotation180,
					}, {
						Tile: &model.Tile{
							Shape:    model.ShapeI,
							Treasure: model.NoTreasure,
						},
						Rotation: model.Rotation180,
					},
				},
			},
		})

		assert.Equal(t, `
───┌─────
 · │B  A 
───│ ┌┐ ┌
│ │┘ │───
│·│ C│ · 
│ │──┘───
│ ││ └───
│·││D  · 
│ │└─────
`[1:], writer.String())
	})
}
