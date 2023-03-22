package model

import (
	"fmt"
	"math"
)

type BoardRowTemplate map[int]*BoardTile
type BoardLineTemplate map[int]BoardRowTemplate
type BoardTemplate map[int]BoardLineTemplate

var (
	boardTemplate BoardTemplate = BoardTemplate{
		3: BoardLineTemplate{
			0: BoardRowTemplate{
				0: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation270,
				},
				2: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation0,
				},
			},
			2: BoardRowTemplate{
				0: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation180,
				},
				2: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation90,
				},
			},
		},
		7: BoardLineTemplate{
			0: BoardRowTemplate{
				0: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation270,
				},
				2: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'A',
					},
					Rotation: Rotation180,
				},
				4: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'B',
					},
					Rotation: Rotation180,
				},
				6: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation0,
				},
			},
			2: BoardRowTemplate{
				0: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'C',
					},
					Rotation: Rotation90,
				},
				2: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'D',
					},
					Rotation: Rotation90,
				},
				4: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'E',
					},
					Rotation: Rotation180,
				},
				6: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'F',
					},
					Rotation: Rotation270,
				},
			},
			4: BoardRowTemplate{
				0: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'G',
					},
					Rotation: Rotation90,
				},
				2: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'H',
					},
					Rotation: Rotation0,
				},
				4: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'I',
					},
					Rotation: Rotation270,
				},
				6: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'J',
					},
					Rotation: Rotation270,
				},
			},
			6: BoardRowTemplate{
				0: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation180,
				},
				2: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'K',
					},
					Rotation: Rotation0,
				},
				4: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeT,
						Treasure: 'L',
					},
					Rotation: Rotation0,
				},
				6: &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation90,
				},
			},
		},
	}
)

var (
	treasureSkipTShaped = map[int]int{
		3: 0,
		7: 12,
	}
)

// generateTiles generates tile list for the given board size.
// It will only generate size*size - 3 cards, since the tiles on each corner is
// predefined (fixed V-shaped).
func generateTiles(size int) (tiles []*Tile, treasures []Treasure) {
	var (
		tileCount = size*size + 1

		// We need to generate 4 less tiles as the corners are V tiles
		generatedTiles               = tileCount - 4
		tShapedThreshold             = int(math.Round(TShapedPercentage * float64(tileCount)))
		vShapedWithTreasureThreshold = tShapedThreshold + int(math.Round(VShapedWithTreasurePercentage*float64(tileCount)))
		iShapedThreshold             = vShapedWithTreasureThreshold + int(math.Round(IShapedPercentage*float64(tileCount)))
	)
	tiles = make([]*Tile, 0, generatedTiles)
	treasures = make([]Treasure, 0, vShapedWithTreasureThreshold)

	var (
		appendTileWithTreasure = func(shape Shape, i int) {
			treasure := 'A' + Treasure(i)
			tiles = append(tiles, &Tile{
				Shape:    shape,
				Treasure: treasure,
			})
			treasures = append(treasures, treasure)
		}
		appendTileWithoutTreasure = func(shape Shape) {
			tiles = append(tiles, &Tile{
				Shape:    shape,
				Treasure: NoTreasure,
			})
		}
	)

	for i := 0; i < generatedTiles; i++ {
		if i < tShapedThreshold {
			appendTileWithTreasure(ShapeT, i)
		} else if i < vShapedWithTreasureThreshold {
			appendTileWithTreasure(ShapeV, i)
		} else if i < iShapedThreshold {
			appendTileWithoutTreasure(ShapeI)
		} else {
			appendTileWithoutTreasure(ShapeV)
		}
	}

	return tiles, treasures
}

func randomRotation() Rotation {
	switch randomGenerator.Int63n(4) {
	case 0:
		return Rotation0
	case 1:
		return Rotation90
	case 2:
		return Rotation90
	default:
		return Rotation270
	}
}

// NewBoard returns a board for the given size.
func NewBoard(size, playerCount int) (*Board, error) {
	if size != 3 && size != 7 {
		return nil, fmt.Errorf("the board size must be either 3 or 7, got: %d", size)
	}

	if playerCount < 1 || playerCount > 4 {
		return nil, fmt.Errorf("the number of players must be between 1 and 4 included, got: %d", playerCount)
	}

	var (
		tiles, treasures = generateTiles(size)
		treasureCount    = len(treasures)
		targetPerPlayer  = treasureCount / playerCount
	)

	// Skip tiles that have been generated statically
	tiles = tiles[treasureSkipTShaped[size]:]

	randomGenerator.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	randomGenerator.Shuffle(len(treasures), func(i, j int) {
		treasures[i], treasures[j] = treasures[j], treasures[i]
	})

	board := &Board{
		Tiles:                make([][]*BoardTile, size),
		Players:              make([]*Player, playerCount),
		RemainingPlayers:     remainingPlayers[:playerCount],
		RemainingPlayerIndex: 0,
		State:                GameStatePlaceTile,
	}

	for i := 0; i < playerCount; i++ {
		board.Players[i] = players[i].Copy()
		treasureOffset := i * targetPerPlayer

		if board.Players[i].Position.Line == -1 {
			board.Players[i].Position.Line = size - 1
		}

		if board.Players[i].Position.Row == -1 {
			board.Players[i].Position.Row = size - 1
		}

		board.Players[i].Targets = treasures[treasureOffset : treasureOffset+targetPerPlayer]
		board.Players[i].Weights = NewBestHintWeights()
	}

	// The tile index is required here to track placed tiles on the board.
	// This is due to the fact that each corner has a predefined V-shaped fixed
	// tile.
	tileIndex := 0
	for line := 0; line < size; line++ {
		board.Tiles[line] = make([]*BoardTile, size)

		for row := 0; row < size; row++ {
			template, ok := boardTemplate[size][line][row]
			if ok {
				board.Tiles[line][row] = template.Copy()
			} else {
				board.Tiles[line][row] = (&BoardTile{
					Tile:     tiles[tileIndex],
					Rotation: randomRotation(),
				}).Copy()
				tileIndex++
			}
		}
	}

	board.RemainingTile = (&BoardTile{
		Tile:     tiles[len(tiles)-1],
		Rotation: Rotation0,
	}).Copy()

	return board, nil
}
